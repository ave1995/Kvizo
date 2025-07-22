package database

import (
	"kvizo-api/internal/repositories"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DatabaseQuizRepository struct {
	database *gorm.DB
}

func NewDatabaseQuizRepository(db *gorm.DB) *DatabaseQuizRepository {
	if db == nil {
		panic("missing db")
	}

	return &DatabaseQuizRepository{database: db}
}

func (r DatabaseQuizRepository) GetByID(id uuid.UUID) (*repositories.Quiz, error) {
	var databaseQuiz databaseQuiz
	err := r.database.First(&databaseQuiz, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return databaseQuiz.ToDomainQuiz(), nil
}

func (r DatabaseQuizRepository) List() ([]*repositories.Quiz, error) {
	var databaseQuizzes []databaseQuiz

	if err := r.database.Find(&databaseQuizzes).Error; err != nil {
		return nil, err
	}

	var quizzes []*repositories.Quiz
	for _, quiz := range databaseQuizzes {
		quizzes = append(quizzes, quiz.ToDomainQuiz())
	}

	return quizzes, nil
}

func (r DatabaseQuizRepository) Create(quiz *repositories.Quiz) error {
	databaseQuiz, err := ToDatabaseQuiz(quiz)
	if err != nil {
		return err
	}

	if err := r.database.Create(databaseQuiz).Error; err != nil {
		return err
	}

	return nil
}
