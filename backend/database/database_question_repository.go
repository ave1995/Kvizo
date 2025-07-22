package database

import (
	"kvizo-api/internal/repositories"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DatabaseQuestionRepository struct {
	database *gorm.DB
	quizRepo *DatabaseQuizRepository
}

func NewDatabaseQuestionRepository(db *gorm.DB, quizRepo *DatabaseQuizRepository) *DatabaseQuestionRepository {
	if db == nil {
		panic("missing db")
	}
	if quizRepo == nil {
		panic("missing quizRepo")
	}

	return &DatabaseQuestionRepository{database: db, quizRepo: quizRepo}
}

func (r DatabaseQuestionRepository) GetByID(id uuid.UUID) (*repositories.Question, error) {
	var databaseQuestion databaseQuestion
	err := r.database.First(&databaseQuestion, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return databaseQuestion.ToDomainQuestion(), nil
}

func (r DatabaseQuestionRepository) ListByQuizID(quizID uuid.UUID) ([]*repositories.Question, error) {
	var questions []*repositories.Question

	err := r.database.Transaction(func(tx *gorm.DB) error {
		_, err := r.quizRepo.GetByID(quizID)
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

func (r DatabaseQuestionRepository) Create(question *repositories.Question) error {
	databaseQuestion, err := ToDatabaseQuestion(question)
	if err != nil {
		return err
	}

	return r.database.Transaction(func(tx *gorm.DB) error {
		_, err := r.quizRepo.GetByID(databaseQuestion.QuizID)
		if err != nil {
			return err
		}

		return tx.Create(databaseQuestion).Error
	})
}
