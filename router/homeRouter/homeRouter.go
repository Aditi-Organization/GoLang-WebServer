package homeRouter

import (
	"GoLang-WebServer/controller/homeController"

	"github.com/gin-gonic/gin"
)

func HomeRouter(rg *gin.RouterGroup) {
	home := rg.Group("/")

	home.GET("/", homeController.Index)

	home.GET("/about", homeController.About)

	home.GET("/contact", homeController.Contact)
}
