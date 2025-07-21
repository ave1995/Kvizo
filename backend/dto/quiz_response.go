package dto

import (
	"kvizo-api/models"
	"time"

	"github.com/google/uuid"
)

type QuizResponse struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func ToResponse(q *models.Quiz) QuizResponse {
	return QuizResponse{
		ID:          q.ID,
		Title:       q.Title,
		Description: q.Description,
		CreatedAt:   q.CreatedAt,
		UpdatedAt:   q.UpdatedAt,
	}
}
