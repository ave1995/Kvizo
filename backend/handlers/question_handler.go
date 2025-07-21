package handlers

import (
	"errors"
	"kvizo-api/database"
	"kvizo-api/dto"
	"kvizo-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CreateQuestionHandler godoc
// @Summary Create a new question
// @Description Create a question with four options under a specific quiz
// @Tags Questions
// @Accept json
// @Produce json
// @Param quiz_id path string true "Quiz ID"
// @Param question body dto.CreateQuestionRequest true "Question info"
// @Success 201 {object} models.Question
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /quizzes/{quiz_id}/questions [post]
func CreateQuestionHandler(c *gin.Context) {
	quizIDStr := c.Param("quiz_id")
	quizID, err := uuid.Parse(quizIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid quiz_id"})
		return
	}

	// Ensure the Quiz exists
	var quiz models.Quiz
	if err := database.DB.First(&quiz, "id = ?", quizID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "quiz not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
		return
	}

	var req dto.CreateQuestionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	question := models.Question{
		QuizID:  quizID,
		Title:   req.Title,
		OptionA: req.OptionA,
		OptionB: req.OptionB,
		OptionC: req.OptionC,
		OptionD: req.OptionD,
		Answer:  req.Answer,
	}

	if err := database.DB.Create(&question).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create question"})
		return
	}

	c.JSON(http.StatusCreated, question)
}
