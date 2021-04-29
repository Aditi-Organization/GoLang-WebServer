package reviewController

import (
	"GoLang-WebServer/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateReviewAPI /api/review/:id?r= <POST>
// UpdateReviewAPI /api/review/:rid?r= <PUT>
// ListReviewsAPI /api/review/ <GET>
// DeleteReviewAPI /api/review/:rid <DELETE>

func CreateReview(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func UpdateReview(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func ListReview(c *gin.Context) {
	movieId := c.Param("id")
	// fmt.Println(movieId)
	result := models.GetMovieReviews(movieId)
	// models.PrintAllReviews()
	fmt.Println(result)
	c.JSON(http.StatusOK, result)
}

func DeleteReview(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
