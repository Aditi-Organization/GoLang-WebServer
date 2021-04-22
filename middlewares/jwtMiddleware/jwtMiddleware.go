package jwtMiddleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

var jwtKey = []byte("SuperStrongPasswordMayBeNot!")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func VerifyJSONWebToken(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	fmt.Println(token)
	tokenString := strings.Split(token, " ")[1]

	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			c.JSON(http.StatusUnauthorized, bson.M{})
		}
		c.JSON(http.StatusBadRequest, bson.M{})
		return
	}
	if !tkn.Valid {
		c.JSON(http.StatusUnauthorized, bson.M{})
		return
	}
	fmt.Println("Successfully verified")
}
