package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AnswerOption string

const (
	OptionA AnswerOption = "A"
	OptionB AnswerOption = "B"
	OptionC AnswerOption = "C"
	OptionD AnswerOption = "D"
)

type Question struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	Title       string       `json:"title"`
	OptionAText string       `json:"option_a"`
	OptionBText string       `json:"option_b"`
	OptionCText string       `json:"option_c"`
	OptionDText string       `json:"option_d"`
	Answer      AnswerOption `gorm:"type:char(1)" json:"answer"`

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
