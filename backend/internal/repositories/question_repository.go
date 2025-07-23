package repositories

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type AnswerOption uint8

const (
	OptionA AnswerOption = iota + 1
	OptionB
	OptionC
	OptionD
)

// TODO: lepší práce s Options
type Question struct {
	ID string

	QuizID string

	Title   string
	OptionA string
	OptionB string
	OptionC string
	OptionD string
	Answer  AnswerOption

	CreatedAt time.Time
	UpdatedAt time.Time
}

type QuestionID uuid.UUID

type QuestionRepository interface {
	GetByID(ctx context.Context, id QuestionID) (*Question, error)
	ListByQuizID(ctx context.Context, quizID QuizID) ([]*Question, error)
	Create(ctx context.Context, quiz *Question) error
	// Update(quiz *Question) error
	// Delete(id uuid.UUID) error
}
