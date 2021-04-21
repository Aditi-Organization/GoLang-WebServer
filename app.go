package main

import (
	// homeRoutes "SPL/routes/homeRoutes"
	"GoLang-WebServer/models"
	"GoLang-WebServer/router/apiRouter"

	"github.com/gin-gonic/gin"
)

func main() {
	// initialize mongodb server
	models.ConnectAndInitialize()

	router := gin.Default()

	api := router.Group("/api/")
	apiRouter.RouteApi(api)
	// homeRouter.HomeRouter(home)

	// movie := router.Group("/movies/")
	// movieRouter.MovieRouter(movie)

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
