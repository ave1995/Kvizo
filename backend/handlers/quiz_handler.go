package handlers

import (
	"kvizo-api/database"
	"kvizo-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateQuizRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}

// CreateQuizHandler godoc
// @Summary Create a new quiz
// @Description Create a quiz with title and description
// @Tags quizzes
// @Accept json
// @Produce json
// @Param quiz body CreateQuizRequest true "Quiz info"
// @Success 201 {object} models.Quiz
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /quizzes [post]
func CreateQuizHandler(c *gin.Context) {
	var req CreateQuizRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	quiz := models.Quiz{
		Title:       req.Title,
		Description: req.Description,
	}

	if err := database.DB.Create(&quiz).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create quiz"})
		return
	}

	c.JSON(http.StatusCreated, quiz)
}
