package dto

type QuizRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}
