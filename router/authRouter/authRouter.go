package authRouter

import (
	"GoLang-WebServer/controller/authController"

	"github.com/gin-gonic/gin"
)

// LoginAPI /api/auth/login
// SignUpAPI /api/auth/signup
// LogoutAPI /api/auth/logout

func AuthRouter(rg *gin.RouterGroup) {
	auth := rg.Group("/")

	auth.GET("/login", authController.Login)

	auth.GET("/signup", authController.SignUp)

	auth.GET("/logout", authController.Logout)
}
