package main

import (
	"ERoBP/controller"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"os"
)

func init() {
	// 设置日志
	//logrus.SetOutput(&lumberjack.Logger{
	//	Filename:   Conf.LogPath,
	//	MaxSize:    100, // MB
	//	MaxBackups: 30,
	//	MaxAge:     0, // Disable age-based rotation
	//	Compress:   true,
	//})

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	//测试
	logrus.SetOutput(os.Stdout)
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func main() {
	router := gin.Default()

	router.Static("/front", "./front")

	router.GET("/", func(c *gin.Context) {
		c.File("./front/index.html")
	})

	user := router.Group("/user")
	{
		user.GET("/register", func(c *gin.Context) {
			c.File("./front/user/register.html")
		})
		user.POST("/register", controller.UserRegister)

		user.GET("/login", func(c *gin.Context) {
			c.File("./front/user/login.html")
		})
		user.POST("/login", controller.UserLogin)

		user.GET("/logout", controller.LogoutUser)

	}

	user.POST("/refresh-token", controller.RefreshJWT)

	router.GET("/home", func(c *gin.Context) {
		c.String(http.StatusOK, "欢迎来到投标专家评审系统的主页!")
	})

	protected := router.Group("/protected")
	protected.Use(controller.JWTMiddleware())
	{
		protected.GET("/verify", controller.VerifyHandler)
	}
	router.Run(":8080")
}
