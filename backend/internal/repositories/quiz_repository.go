package repositories

import (
	"time"

	"github.com/google/uuid"
)

type Quiz struct {
	ID string

	Title       string
	Description string

	Questions []*Question

	CreatedAt time.Time
	UpdatedAt time.Time
}

type QuizRepository interface {
	GetByID(id uuid.UUID) (*Quiz, error)
	List() ([]*Quiz, error)
	Create(quiz *Quiz) error
	// Update(quiz *Quiz) error
	// Delete(id uuid.UUID) error
}
