package dto

type CreateQuizRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}
