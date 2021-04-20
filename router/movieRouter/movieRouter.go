package movieRouter

import (
	"GoLang-WebServer/controller/movieController"

	"github.com/gin-gonic/gin"
)

func MovieRouter(rg *gin.RouterGroup) {
	movie := rg.Group("/")

	movie.GET("/", movieController.Index)

	movie.GET("/:id", movieController.MovieDetails)

}
