package handlers

import (
	"kvizo-api/dto"
	"kvizo-api/internal/repositories"
	"kvizo-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type QuizHandler struct {
	service *services.QuizService
}

func NewQuizHandler(s *services.QuizService) *QuizHandler {
	return &QuizHandler{service: s}
}

// CreateQuizHandler godoc
// @Summary Create a new quiz
// @Description Create a quiz with title and description
// @Tags quizzes
// @Accept json
// @Produce json
// @Param quiz body dto.CreateQuizRequest true "Quiz info"
// @Success 201 {object} repositories.Quiz
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /quizzes [post]
func (h *QuizHandler) CreateQuizHandler(c *gin.Context) {
	var req dto.CreateQuizRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	quiz := repositories.Quiz{
		Title:       req.Title,
		Description: req.Description,
	}

	if err := h.service.Create(&quiz); err != nil {
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
func (h *QuizHandler) GetQuizzesHandler(c *gin.Context) {
	quizzes, _ := h.service.List()

	//TODO: dvakrát foreach s backend/database/database_quiz_repository.go
	var responses []dto.QuizResponse
	for _, quiz := range quizzes {
		responses = append(responses, dto.ToResponse(quiz))
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
func (h *QuizHandler) GetQuizHandler(c *gin.Context) {
	idParam := c.Param("id")
	quiz, err := h.service.GetByID(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid quiz ID"})
		return
	}

	c.JSON(http.StatusOK, dto.ToResponse(quiz))
}
