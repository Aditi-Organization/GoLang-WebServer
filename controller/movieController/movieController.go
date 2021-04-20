package movieController

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "movie/index.tmpl", gin.H{})
}

func MovieDetails(c *gin.Context) {
	id := c.Param("id")
	c.HTML(http.StatusOK, "movie/movie.tmpl", gin.H{"id": id})
}
