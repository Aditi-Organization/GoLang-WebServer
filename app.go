package main

import (
	// homeRoutes "SPL/routes/homeRoutes"
	"GoLang-WebServer/models"
	"GoLang-WebServer/router/apiRouter"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// initialize mongodb server
	models.ConnectAndInitialize()

	router := gin.Default()

	// using default CORS until deployed
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))

	api := router.Group("/api")
	apiRouter.RouteApi(api)

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
