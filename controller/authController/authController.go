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
	c.JSON(http.StatusOK, gin.H{})
}

func SignUp(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
