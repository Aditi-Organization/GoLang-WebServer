package main

import (
	// homeRoutes "SPL/routes/homeRoutes"
	"GoLang-WebServer/router/homeRouter"
	"GoLang-WebServer/router/movieRouter"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/**/*")

	home := router.Group("/")
	homeRouter.HomeRouter(home)

	movie := router.Group("/movie")
	movieRouter.MovieRouter(movie)

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
