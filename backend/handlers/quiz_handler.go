package handlers

import (
	"kvizo-api/dto"
	"kvizo-api/internal/repositories"
	"kvizo-api/internal/responses"
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
// @Param quiz body dto.QuizRequest true "Quiz info"
// @Success 201 {object} repositories.Quiz
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/quizzes [post]
func (h *QuizHandler) CreateQuizHandler(c *gin.Context) {
	var req dto.QuizRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		responses.RespondWithInternalError(c, err)
		return
	}

	quiz := repositories.Quiz{
		Title:       req.Title,
		Description: req.Description,
	}

	if err := h.service.Create(c, &quiz); err != nil {
		responses.RespondWithInternalError(c, err)
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
// @Router /api/quizzes [get]
func (h *QuizHandler) GetQuizzesHandler(c *gin.Context) {
	quizzes, err := h.service.List(c)
	if err != nil {
		responses.RespondWithInternalError(c, err)
		return
	}

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
// @Router /api/quiz/{id} [get]
func (h *QuizHandler) GetQuizHandler(c *gin.Context) {
	idParam := c.Param("id")
	quiz, err := h.service.GetQuiz(c, idParam)
	if err != nil {
		responses.RespondWithInternalError(c, err)
		return
	}

	c.JSON(http.StatusOK, dto.ToResponse(quiz))
}

// UpdateQuizHandler godoc
// @Summary Update a quiz
// @Description Update the title and description of an existing quiz
// @Tags quiz
// @Accept json
// @Produce json
// @Param id path string true "Quiz ID"
// @Param quiz body dto.QuizRequest true "Updated quiz info"
// @Success 200 {object} repositories.Quiz
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/quiz/{id} [put]
func (h *QuizHandler) UpdateQuizHandler(c *gin.Context) {
	var req dto.QuizRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		responses.RespondWithInternalError(c, err)
		return
	}

	quizID := c.Param("id")

	quiz := repositories.Quiz{
		ID:          quizID,
		Title:       req.Title,
		Description: req.Description,
	}

	if err := h.service.Update(c, &quiz); err != nil {
		responses.RespondWithInternalError(c, err)
		return
	}

	c.JSON(http.StatusOK, quiz)
}

// DeleteQuizHandler deletes a quiz by its ID.
// @Summary Delete a quiz
// @Description Deletes a quiz by ID.
// @Tags quiz
// @Accept json
// @Produce json
// @Param id path string true "Quiz ID"
// @Success 200 {object} map[string]string "message: Quiz deleted successfully"
// @Failure 500 {object} map[string]string "error message"
// @Router /api/quiz/{id} [delete]
func (h *QuizHandler) DeleteQuizHandler(c *gin.Context) {
	idParam := c.Param("id")
	err := h.service.Delete(c, idParam)
	if err != nil {
		responses.RespondWithInternalError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Quiz deleted successfully",
	})
}
