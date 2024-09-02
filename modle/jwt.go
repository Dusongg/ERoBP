package modle

import "github.com/dgrijalva/jwt-go"

var JWTSecret = []byte("your_secret_key")

type Claims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}
