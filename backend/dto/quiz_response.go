package dto

import (
	"kvizo-api/internal/repositories"
	"time"
)

type QuizResponse struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func ToResponse(q *repositories.Quiz) QuizResponse {
	return QuizResponse{
		ID:          q.ID,
		Title:       q.Title,
		Description: q.Description,
		CreatedAt:   q.CreatedAt,
		UpdatedAt:   q.UpdatedAt,
	}
}
