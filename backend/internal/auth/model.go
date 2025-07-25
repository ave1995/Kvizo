package auth

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DatabaseUser struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey"`
	Email        string    `gorm:"uniqueIndex;not null"`
	PasswordHash string    `gorm:"not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (dbUser *DatabaseUser) BeforeCreate(tx *gorm.DB) (err error) {
	if dbUser.ID == uuid.Nil {
		dbUser.ID = uuid.New()
	}
	return
}

func (dbUser *DatabaseUser) toUser() *User {
	return &User{
		ID:    dbUser.ID,
		Email: dbUser.Email,
	}
}

type User struct {
	ID    uuid.UUID
	Email string
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token     string `json:"token"`
	ExpiresIn int64  `json:"expiresIn"`
}
