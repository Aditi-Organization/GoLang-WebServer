package reviewController

import (
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
	c.JSON(http.StatusOK, gin.H{})
}

func DeleteReview(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
