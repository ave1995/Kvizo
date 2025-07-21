package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Quiz struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	Title       string `json:"title"`
	Description string `json:"description,omitempty"`

	Questions []Question `json:"questions"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (q *Quiz) BeforeCreate(tx *gorm.DB) (err error) {
	if q.ID == uuid.Nil {
		q.ID = uuid.New()
	}
	return
}
