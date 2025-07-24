package handlers

import (
	"kvizo-api/dto"
	"kvizo-api/internal/repositories"
	"kvizo-api/internal/responses"
	"kvizo-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type QuestionHandler struct {
	service *services.QuestionService
}

func NewQuestionHandler(s *services.QuestionService) *QuestionHandler {
	return &QuestionHandler{service: s}
}

// CreateQuestionHandler godoc
// @Summary Create a new question
// @Description Create a question with four options under a specific quiz
// @Tags questions
// @Accept json
// @Produce json
// @Param quiz_id path string true "Quiz ID"
// @Param question body dto.CreateQuestionRequest true "Question info"
// @Success 201 {object} repositories.Question
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /quizzes/{quiz_id}/questions [post]
func (h *QuestionHandler) CreateQuestionHandler(c *gin.Context) {
	var req dto.CreateQuestionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		responses.RespondWithInternalError(c, err)
		return
	}

	quizID := c.Param("quiz_id")

	question := repositories.Question{
		QuizID:  quizID,
		Title:   req.Title,
		OptionA: req.OptionA,
		OptionB: req.OptionB,
		OptionC: req.OptionC,
		OptionD: req.OptionD,
		Answer:  req.Answer,
	}

	if err := h.service.Create(c, &question); err != nil {
		responses.RespondWithInternalError(c, err)
		return
	}

	c.JSON(http.StatusCreated, question)
}

// GetQuestionsForQuizHandler godoc
// @Summary Get all questions for a quiz
// @Description Retrieve all questions belonging to a specific quiz
// @Tags questions
// @Produce json
// @Param quiz_id path string true "Quiz ID"
// @Success 200 {array} repositories.Question
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /quizzes/{quiz_id}/questions [get]
func (h *QuestionHandler) GetQuestionsForQuizHandler(c *gin.Context) {
	idParam := c.Param("quiz_id")
	questions, err := h.service.ListByQuizID(c, idParam)
	if err != nil {
		responses.RespondWithInternalError(c, err)
		return
	}

	c.JSON(http.StatusOK, questions)
}

// DeleteQuestionHandler deletes a quiz by its ID.
// @Summary Delete a question
// @Description Deletes a question by ID.
// @Tags question
// @Accept json
// @Produce json
// @Param id path string true "Question ID"
// @Success 200 {object} map[string]string "message: Question deleted successfully"
// @Failure 500 {object} map[string]string "error message"
// @Router /questions/{id} [delete]
func (h *QuestionHandler) DeleteQuestionHandler(c *gin.Context) {
	idParam := c.Param("id")
	err := h.service.Delete(c, idParam)
	if err != nil {
		responses.RespondWithInternalError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Question deleted successfully",
	})
}
