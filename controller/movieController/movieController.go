package movieController

import (
	"GoLang-WebServer/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ListMoviesAPI /api/movies/ <GET>
// SearchMovieAPI /api/movies?q=<movieName> <POST>
// DisplayMovieAPI /api/movies/:id <GET>

func Index(c *gin.Context) {
	movieList := models.FindAll()
	// fmt.Print(movieList)
	// for i := 0; i < len(movieList); i++ {
	// 	fmt.Println(movieList[i])
	// }
	c.JSON(http.StatusOK, movieList)
}

func DisplayMovie(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func SearchMovie(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"id": id})
}
