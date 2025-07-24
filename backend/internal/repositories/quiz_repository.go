package repositories

import (
	"context"
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

type QuizID uuid.UUID

type QuizRepository interface {
	GetByID(ctx context.Context, id QuizID) (*Quiz, error)
	List(ctx context.Context) ([]*Quiz, error)
	Create(ctx context.Context, quiz *Quiz) error
	Update(ctx context.Context, quiz *Quiz) error
	Delete(ctx context.Context, id QuizID) error
}
