package movieRouter

import (
	"GoLang-WebServer/controller/movieController"

	"github.com/gin-gonic/gin"
)

// ListMoviesAPI /api/movies/ <GET>
// SearchMovieAPI /api/movies?q=<movieName> <POST>
// DisplayMovieAPI /api/movies/:id <GET>

func MovieRouter(rg *gin.RouterGroup) {
	movie := rg.Group("/")

	movie.GET("/", movieController.Index)

	movie.POST("/", movieController.SearchMovie)

	movie.GET("/:id", movieController.DisplayMovie)

}
