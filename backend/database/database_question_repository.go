package database

import (
	"context"
	"kvizo-api/internal/repositories"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DatabaseQuestionRepository struct {
	gorm     *gorm.DB
	quizRepo *DatabaseQuizRepository
}

func NewDatabaseQuestionRepository(db *gorm.DB, quizRepo *DatabaseQuizRepository) *DatabaseQuestionRepository {
	if db == nil {
		panic("missing db")
	}
	if quizRepo == nil {
		panic("missing quizRepo")
	}

	return &DatabaseQuestionRepository{gorm: db, quizRepo: quizRepo}
}

func (r *DatabaseQuestionRepository) GetByID(ctx context.Context, id repositories.QuestionID) (*repositories.Question, error) {
	result, err := getByID[databaseQuestion](r.gorm.WithContext(ctx), uuid.UUID(id))
	if err != nil {
		return nil, err
	}

	return result.ToDomainQuestion(), nil
}

func (r *DatabaseQuestionRepository) ListByQuizID(ctx context.Context, quizID repositories.QuizID) ([]*repositories.Question, error) {
	var questions []*repositories.Question

	err := r.gorm.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		_, err := getByID[databaseQuiz](tx, uuid.UUID(quizID))
		if err != nil {
			return err
		}

		var dbQuestions []databaseQuestion
		if err := tx.Where("quiz_id = ?", quizID).Find(&dbQuestions).Error; err != nil {
			return err
		}

		for _, q := range dbQuestions {
			questions = append(questions, q.ToDomainQuestion())
		}

		return nil
	})

	return questions, err
}

func (r *DatabaseQuestionRepository) Create(ctx context.Context, question *repositories.Question) error {
	databaseQuestion, err := ToDatabaseQuestion(question)
	if err != nil {
		return err
	}

	return r.gorm.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		_, err := getByID[databaseQuiz](tx, databaseQuestion.QuizID)
		if err != nil {
			return err
		}

		return tx.Create(databaseQuestion).Error
	})
}
