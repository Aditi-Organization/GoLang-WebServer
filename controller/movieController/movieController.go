package movieController

import (
	"GoLang-WebServer/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	movieList := models.FindAll()
	fmt.Print(movieList)
	// for i := 0; i < len(movieList); i++ {
	// 	fmt.Println(movieList[i])
	// }
	c.HTML(http.StatusOK, "movie/index.tmpl", movieList)
}

func MovieDetails(c *gin.Context) {
	id := c.Param("id")
	c.HTML(http.StatusOK, "movie/movie.tmpl", gin.H{"id": id})
}
