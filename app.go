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

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		// AllowOriginFunc: func(origin string) bool {
		// 	return origin == "https://github.com"
		// },
		// MaxAge: 12 * time.Hour,
	}))

	api := router.Group("/api/")
	apiRouter.RouteApi(api)
	// homeRouter.HomeRouter(home)

	// movie := router.Group("/movies/")
	// movieRouter.MovieRouter(movie)

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
