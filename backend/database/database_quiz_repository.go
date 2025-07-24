package database

import (
	"context"
	"kvizo-api/internal/repositories"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type QuizRepository struct {
	gorm *gorm.DB
}

func NewDatabaseQuizRepository(db *gorm.DB) *QuizRepository {
	if db == nil {
		panic("missing db")
	}

	return &QuizRepository{gorm: db}
}

func (r *QuizRepository) GetByID(ctx context.Context, id repositories.QuizID) (*repositories.Quiz, error) {
	result, err := getByID[databaseQuiz](r.gorm.WithContext(ctx), uuid.UUID(id))
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
	databaseQuiz, err := toDatabaseQuiz(quiz)
	if err != nil {
		return err
	}

	if err := r.gorm.WithContext(ctx).Create(databaseQuiz).Error; err != nil {
		return err
	}

	return nil
}
