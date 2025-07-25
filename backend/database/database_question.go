package database

import (
	"kvizo-api/internal/repositories"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type databaseQuestion struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	QuizID uuid.UUID `gorm:"type:uuid;not null;index"` // FK

	Title   string
	OptionA string
	OptionB string
	OptionC string
	OptionD string
	Answer  uint8 `gorm:"type:smallint"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

const QuestionTableName = "questions"

func (*databaseQuestion) TableName() string {
	return QuestionTableName
}

func (q *databaseQuestion) BeforeCreate(tx *gorm.DB) (err error) {
	if q.ID == uuid.Nil {
		q.ID = uuid.New()
	}
	return
}

func (g *databaseQuestion) ToDomainQuestion() *repositories.Question {
	return &repositories.Question{
		ID:        g.ID.String(),
		QuizID:    g.QuizID.String(),
		Title:     g.Title,
		OptionA:   g.OptionA,
		OptionB:   g.OptionB,
		OptionC:   g.OptionC,
		OptionD:   g.OptionD,
		Answer:    repositories.AnswerOption(g.Answer),
		CreatedAt: g.CreatedAt,
		UpdatedAt: g.UpdatedAt,
	}
}

func ToDatabaseQuestion(q *repositories.Question, requireID bool) (*databaseQuestion, error) {
	quizID, err := parseRequiredUUID("QuizID", q.QuizID)
	if err != nil {
		return nil, err
	}

	var id uuid.UUID
	if requireID {
		id, err = parseRequiredUUID("ID", q.ID)
		if err != nil {
			return nil, err
		}
	}

	return &databaseQuestion{
		ID:      id,
		QuizID:  quizID,
		Title:   q.Title,
		OptionA: q.OptionA,
		OptionB: q.OptionB,
		OptionC: q.OptionC,
		OptionD: q.OptionD,
		Answer:  uint8(q.Answer),
	}, nil
}

func ToDomainQuestions(questions []databaseQuestion) []*repositories.Question {
	result := make([]*repositories.Question, 0, len(questions))
	for i, q := range questions {
		result[i] = q.ToDomainQuestion()
	}
	return result
}
