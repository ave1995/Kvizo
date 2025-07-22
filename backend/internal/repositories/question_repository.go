package repositories

import (
	"time"

	"github.com/google/uuid"
)

type AnswerOption uint8

const (
	OptionA AnswerOption = 1
	OptionB AnswerOption = 2
	OptionC AnswerOption = 3
	OptionD AnswerOption = 4
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

type QuestionRepository interface {
	GetByID(id uuid.UUID) (*Question, error)
	ListByQuizID(quizID uuid.UUID) ([]*Question, error)
	Create(quiz *Question) error
	// Update(quiz *Question) error
	// Delete(id uuid.UUID) error
}
