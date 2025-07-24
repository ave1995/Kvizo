package auth

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthenticationRepository interface {
	RegisterUser(ctx context.Context, email, password string) (*User, error)
	AuthenticateUser(ctx context.Context, email, password string) (*User, error)
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
