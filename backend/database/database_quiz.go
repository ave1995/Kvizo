package database

import (
	"kvizo-api/internal/repositories"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type databaseQuiz struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	Title       string
	Description string

	Questions []databaseQuestion `gorm:"foreignKey:QuizID;references:ID"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (databaseQuiz) TableName() string {
	return "quizzes"
}

func (q *databaseQuiz) BeforeCreate(tx *gorm.DB) (err error) {
	if q.ID == uuid.Nil {
		q.ID = uuid.New()
	}
	return
}

func (g *databaseQuiz) ToDomainQuiz() *repositories.Quiz {
	return &repositories.Quiz{
		ID:          g.ID.String(),
		Title:       g.Title,
		Description: g.Description,
		Questions:   ToDomainQuestions(g.Questions),
		CreatedAt:   g.CreatedAt,
		UpdatedAt:   g.UpdatedAt,
	}
}

func ToDatabaseQuiz(dq *repositories.Quiz) (*databaseQuiz, error) {
	return &databaseQuiz{
		Title:       dq.Title,
		Description: dq.Description,
		CreatedAt:   dq.CreatedAt,
		UpdatedAt:   dq.UpdatedAt,
	}, nil
}
