package authController

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtKey = []byte("SuperStrongPasswordMayBeNot!")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func SignJWT(c *gin.Context, userName string, password string) (string, int) {
	// 5 minutes
	expirationTime := time.Now().Add(60 * time.Minute)

	claims := &Claims{
		Username: userName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println("Token string", tokenString)
	return tokenString, int(expirationTime.Unix())
}
