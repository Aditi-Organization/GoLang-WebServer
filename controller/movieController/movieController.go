package movieController

import (
	"GoLang-WebServer/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ListMoviesAPI /api/movies/ <GET>
// SearchMovieAPI /api/movies?q=<movieName> <POST>
// DisplayMovieAPI /api/movies/:id <GET>

func Index(c *gin.Context) {
	movieList := models.FindAllVersion2()
	fmt.Print(movieList)
	c.JSON(http.StatusOK, movieList)
}

func DisplayMovie(c *gin.Context) {
	//fmt.Println(c.Params)
	movieId := c.Param("id")
	fmt.Println(movieId)
	if movieId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Provide ID"})
		return
	}
	movieObject := models.ShowMovie(movieId)
	fmt.Println(movieObject)
	// c.JSON(http.StatusOK, gin.H{})
	c.JSON(http.StatusOK, gin.H{"name": movieObject.Name, "_id": movieObject.Id})
}

func SearchMovie(c *gin.Context) {
	movieName := c.Request.FormValue("movieName")
	if movieName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Please fill the field: movie name"})
		return
	}
	movieObject := models.FindMovie(movieName)
	fmt.Println(movieObject)
	//if (!movieObject)  {
	//	c.JSON(http.StatusNotFound, gin.H{"Not found": "Movie is not available"})
	//}

	c.JSON(http.StatusOK, gin.H{"name": movieObject.Name, "_id": movieObject.Id})
}
