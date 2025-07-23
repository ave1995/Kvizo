package database

import (
	"context"
	"kvizo-api/internal/repositories"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DatabaseQuizRepository struct {
	gorm *gorm.DB
}

func NewDatabaseQuizRepository(db *gorm.DB) *DatabaseQuizRepository {
	if db == nil {
		panic("missing db")
	}

	return &DatabaseQuizRepository{gorm: db}
}

func (r *DatabaseQuizRepository) GetByID(ctx context.Context, id repositories.QuizID) (*repositories.Quiz, error) {
	result, err := getByID[databaseQuiz](r.gorm.WithContext(ctx), uuid.UUID(id))
	if err != nil {
		return nil, err
	}

	return result.ToDomainQuiz(), nil
}

func (r *DatabaseQuizRepository) List(ctx context.Context) ([]*repositories.Quiz, error) {
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

func (r *DatabaseQuizRepository) Create(ctx context.Context, quiz *repositories.Quiz) error {
	databaseQuiz, err := ToDatabaseQuiz(quiz)
	if err != nil {
		return err
	}

	if err := r.gorm.WithContext(ctx).Create(databaseQuiz).Error; err != nil {
		return err
	}

	return nil
}
