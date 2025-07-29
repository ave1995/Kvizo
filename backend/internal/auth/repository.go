package auth

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthRepository interface {
	RegisterUser(ctx context.Context, email, password string) (*User, error)
	AuthenticateUser(ctx context.Context, email, password string) (*User, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (*User, error)

	SaveRefreshToken(ctx context.Context, token RefreshToken) error
	RevokeRefreshToken(ctx context.Context, token string) error
	GetRefreshToken(ctx context.Context, token string) (*RefreshToken, error)
}

type DatabaseUserRepository struct {
	gorm *gorm.DB
}

func NewDatabaseUserRepository(gorm *gorm.DB) (*DatabaseUserRepository, error) {
	if gorm == nil {
		return nil, errors.New("NewDatabaseUserRepository: missing gorm db")
	}

	return &DatabaseUserRepository{gorm: gorm}, nil
}

func getByEmail(gorm *gorm.DB, email string) (*DatabaseUser, error) {
	var user *DatabaseUser

	if err := gorm.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func getByID(gorm *gorm.DB, id uuid.UUID) (*DatabaseUser, error) {
	var user *DatabaseUser

	if err := gorm.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *DatabaseUserRepository) RegisterUser(ctx context.Context, email, password string) (*User, error) {
	var user *User

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	err = r.gorm.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if _, err := getByEmail(tx, email); err == nil {
			return err
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			// An unexpected error occurred
			return err
		}

		dbUser := &DatabaseUser{
			Email:        email,
			PasswordHash: string(hashedPassword),
		}

		err = tx.Create(dbUser).Error
		if err != nil {
			return err
		}

		user = dbUser.toUser()

		return nil
	})

	return user, err
}

func (r *DatabaseUserRepository) AuthenticateUser(ctx context.Context, email, password string) (*User, error) {
	dbUser, err := getByEmail(r.gorm.WithContext(ctx), email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbUser.PasswordHash), []byte(password))
	if err != nil {
		return nil, err
	}

	return dbUser.toUser(), nil
}

func (r *DatabaseUserRepository) GetUserByID(ctx context.Context, id uuid.UUID) (*User, error) {
	dbUser, err := getByID(r.gorm.WithContext(ctx), id)
	if err != nil {
		return nil, err
	}

	return dbUser.toUser(), nil
}

func (r *DatabaseUserRepository) SaveRefreshToken(ctx context.Context, token RefreshToken) error {
	dbToken := token.toDatabaseRefreshToken()

	return r.gorm.WithContext(ctx).Create(&dbToken).Error
}

func (r *DatabaseUserRepository) RevokeRefreshToken(ctx context.Context, token string) error {
	return r.gorm.WithContext(ctx).
		Model(&DatabaseRefreshToken{}).
		Where("token = ?", token).
		Update("revoked", true).Error
}

func (r *DatabaseUserRepository) GetRefreshToken(ctx context.Context, token string) (*RefreshToken, error) {
	var dbToken DatabaseRefreshToken

	err := r.gorm.WithContext(ctx).
		Where("token = ?", token).
		First(&dbToken).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	result := dbToken.toRefreshToken()
	return &result, nil
}
