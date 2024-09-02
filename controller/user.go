package controller

import (
	"ERoBP/common"
	"ERoBP/config"
	"ERoBP/modle"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func UserLogin(c *gin.Context) {
	var inputInfo User
	if err := c.ShouldBindJSON(&inputInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	var userinfo modle.UserInfo
	if err := modle.Db.Where("name = ?", inputInfo.Username).First(&userinfo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userinfo.Password), []byte(inputInfo.Password)); err == nil {
		accessToken, err := common.GenerateJWT(userinfo.ID, userinfo.Name, userinfo.Role, 15*time.Minute)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate access token"})
			return
		}
		fmt.Println(userinfo.ID, userinfo.Name, userinfo.Role)

		refreshToken, err := common.GenerateJWT(userinfo.ID, userinfo.Name, userinfo.Role, 7*24*time.Hour)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate refresh token"})
			return
		}

		c.SetCookie("refreshToken", refreshToken, 7*24*60*60, "/", "", false, true)
		c.JSON(http.StatusOK, gin.H{
			"accessToken": accessToken,
			"success":     true,
			"avatarUrl":   userinfo.AvatarUrl})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
	}
}

func UserRegister(c *gin.Context) {
	username := c.PostForm("username")
	phone := c.PostForm("phone")
	password := c.PostForm("password")
	role := c.PostForm("role")

	if modle.Db.Where("phone = ?", phone).First(&modle.UserInfo{}).RowsAffected == 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user already exist"})
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	userinfo := &modle.UserInfo{
		Name:      username,
		Phone:     phone,
		Password:  string(hashedPassword),
		Role:      role,
		AvatarUrl: config.DefaultAvatar,
	}
	if err := modle.Db.Create(userinfo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	//fmt.Println(username, password, email, role)
	c.Redirect(http.StatusSeeOther, "/user/login")
}

func LogoutUser(c *gin.Context) {
	c.SetCookie("refreshToken", "", -1, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}
