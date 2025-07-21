package handlers

import (
	"kvizo-api/database"
	"kvizo-api/dto"
	"kvizo-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateQuizHandler godoc
// @Summary Create a new quiz
// @Description Create a quiz with title and description
// @Tags quizzes
// @Accept json
// @Produce json
// @Param quiz body dto.CreateQuizRequest true "Quiz info"
// @Success 201 {object} models.Quiz
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /quizzes [post]
func CreateQuizHandler(c *gin.Context) {
	var req dto.CreateQuizRequest

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

// GetQuizzesHandler godoc
// @Summary Get all quizzes
// @Description Retrieve all quizzes
// @Tags quizzes
// @Produce json
// @Success 200 {array} dto.QuizResponse
// @Failure 500 {object} map[string]string
// @Router /quizzes [get]
func GetQuizzesHandler(c *gin.Context) {
	var quizzes []models.Quiz

	if err := database.DB.Find(&quizzes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve quizzes"})
		return
	}

	var responses []dto.QuizResponse
	for _, quiz := range quizzes {
		responses = append(responses, dto.ToResponse(&quiz))
	}

	c.JSON(http.StatusOK, responses)
}

// GetQuizHandler godoc
// @Summary Get a quiz by ID
// @Description Retrieve a single quiz by its ID
// @Tags quiz
// @Produce json
// @Param id path string true "Quiz ID"
// @Success 200 {object} dto.QuizResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /quiz/{id} [get]
func GetQuizHandler(c *gin.Context) {
	idParam := c.Param("id")
	quizID, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid quiz ID"})
		return
	}

	var quiz models.Quiz
	if err := database.DB.First(&quiz, "id = ?", quizID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Quiz not found"})
		return
	}

	c.JSON(http.StatusOK, dto.ToResponse(&quiz))
}
