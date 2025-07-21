package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AnswerOption uint8

const (
	OptionA AnswerOption = 1
	OptionB AnswerOption = 2
	OptionC AnswerOption = 3
	OptionD AnswerOption = 4
)

type Question struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	QuizID uuid.UUID `gorm:"type:uuid;not null;index" json:"quiz_id"` // Foreign key

	Title   string       `json:"title"`
	OptionA string       `json:"option_a"`
	OptionB string       `json:"option_b"`
	OptionC string       `json:"option_c"`
	OptionD string       `json:"option_d"`
	Answer  AnswerOption `gorm:"type:smallint" json:"answer"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (q *Question) BeforeCreate(tx *gorm.DB) (err error) {
	if q.ID == uuid.Nil {
		q.ID = uuid.New()
	}
	return
}
