package auth

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DatabaseRefreshToken struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID    uuid.UUID `gorm:"type:uuid;index"`
	Token     uuid.UUID `gorm:"type:uuid;uniqueIndex"`
	ExpiresAt time.Time
	Revoked   bool
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (dbToken DatabaseRefreshToken) toRefreshToken() RefreshToken {
	return RefreshToken{
		ID:        dbToken.ID,
		UserID:    dbToken.UserID,
		Token:     dbToken.Token,
		ExpiresAt: dbToken.ExpiresAt,
		Revoked:   dbToken.Revoked,
		CreatedAt: dbToken.CreatedAt,
	}
}

type RefreshToken struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Token     uuid.UUID
	ExpiresAt time.Time
	Revoked   bool
	CreatedAt time.Time
}

func (token RefreshToken) toDatabaseRefreshToken() DatabaseRefreshToken {
	return DatabaseRefreshToken{
		ID:        token.ID,
		UserID:    token.UserID,
		Token:     token.Token,
		ExpiresAt: token.ExpiresAt,
		Revoked:   token.Revoked,
		CreatedAt: token.CreatedAt,
	}
}

type AccessToken struct {
	Token     string
	ExpiresAt time.Time
	UserID    uuid.UUID
	Email     string
	// Role      string
}
