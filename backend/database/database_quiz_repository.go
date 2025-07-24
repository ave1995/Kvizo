package database

import (
	"context"
	"errors"
	"kvizo-api/internal/repositories"

	"gorm.io/gorm"
)

type QuizRepository struct {
	gorm *gorm.DB
}

func NewDatabaseQuizRepository(db *gorm.DB) (*QuizRepository, error) {
	if db == nil {
		return nil, errors.New("NewDatabaseQuizRepository: missing gorm DB")

	}

	return &QuizRepository{gorm: db}, nil
}

func (r *QuizRepository) GetByID(ctx context.Context, id repositories.QuizID) (*repositories.Quiz, error) {
	result, err := getByID[databaseQuiz](r.gorm.WithContext(ctx), toUUID(id))
	if err != nil {
		return nil, err
	}

	return result.ToDomainQuiz(), nil
}

func (r *QuizRepository) List(ctx context.Context) ([]*repositories.Quiz, error) {
	var databaseQuizzes []databaseQuiz

	if err := r.gorm.WithContext(ctx).Find(&databaseQuizzes).Error; err != nil {
		return nil, err
	}

	var quizzes []*repositories.Quiz
	for _, quiz := range databaseQuizzes {
		quizzes = append(quizzes, quiz.ToDomainQuiz())
	}

	return quizzes, nil
}

func (r *QuizRepository) Create(ctx context.Context, quiz *repositories.Quiz) error {
	dbQuiz, err := toDatabaseQuiz(quiz, false)
	if err != nil {
		return err
	}

	if err := r.gorm.WithContext(ctx).Create(dbQuiz).Error; err != nil {
		return err
	}

	return nil
}

func (r *QuizRepository) Update(ctx context.Context, quiz *repositories.Quiz) error {
	dbQuiz, err := toDatabaseQuiz(quiz, true)
	if err != nil {
		return err
	}

	return r.gorm.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		existing, err := getByID[databaseQuiz](tx, dbQuiz.ID)
		if err != nil {
			return err
		}

		existing.Title = dbQuiz.Title
		existing.Description = dbQuiz.Description

		return tx.Save(existing).Error
	})
}

func (r *QuizRepository) Delete(ctx context.Context, id repositories.QuizID) error {
	return r.gorm.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&databaseQuiz{}, toUUID(id)).Error; err != nil {
			return err
		}
		return nil
	})
}
