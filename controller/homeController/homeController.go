package homeController

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// only names that begin with caps are exported!
func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "home/index.tmpl", gin.H{})
}

func About(c *gin.Context) {
	c.JSON(http.StatusOK, "About")
}

func Contact(c *gin.Context) {
	c.HTML(http.StatusOK, "home/contact.tmpl", gin.H{})
}
