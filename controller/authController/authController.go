package authController

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// only names that begin with caps are exported!

// LoginAPI /api/auth/login
// SignUpAPI /api/auth/signup
// LogoutAPI /api/auth/logout

func Login(c *gin.Context) {
	// c.HTML(http.StatusOK, "home/index.tmpl", gin.H{})
	userName := c.Request.FormValue("username")
	password := c.Request.FormValue("password")
	verifyLogin(userName, password)
	tokenString, expirationTime := SignJWT(c, userName, password)
	// c.SetCookie("token", tokenString, expirationTime, "/", "localhost", true, false)

	c.JSON(http.StatusOK, gin.H{"token": tokenString, "expirationTime": expirationTime})
}

func SignUp(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func verifyLogin(userName string, password string) {

}
