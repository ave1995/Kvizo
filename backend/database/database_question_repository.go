package database

import (
	"context"
	"errors"
	"kvizo-api/internal/repositories"

	"gorm.io/gorm"
)

type QuestionRepository struct {
	gorm     *gorm.DB
	quizRepo *QuizRepository
}

func NewDatabaseQuestionRepository(db *gorm.DB, quizRepo *QuizRepository) (*QuestionRepository, error) {
	if db == nil {
		return nil, errors.New("NewDatabaseQuestionRepository: missing gorm DB")
	}
	if quizRepo == nil {
		return nil, errors.New("NewDatabaseQuestionRepository: missing Quiz Repository")
	}

	return &QuestionRepository{gorm: db, quizRepo: quizRepo}, nil
}

func (r *QuestionRepository) GetByID(ctx context.Context, id repositories.QuestionID) (*repositories.Question, error) {
	result, err := getByID[databaseQuestion](r.gorm.WithContext(ctx), toUUID(id))
	if err != nil {
		return nil, err
	}

	return result.ToDomainQuestion(), nil
}

func (r *QuestionRepository) ListByQuizID(ctx context.Context, quizID repositories.QuizID) ([]*repositories.Question, error) {
	var questions []*repositories.Question

	err := r.gorm.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		_, err := getByID[databaseQuiz](tx, toUUID(quizID))
		if err != nil {
			return err
		}

		var dbQuestions []databaseQuestion
		if err := tx.Where("quiz_id = ?", toUUID(quizID)).Find(&dbQuestions).Error; err != nil {
			return err
		}

		for _, q := range dbQuestions {
			questions = append(questions, q.ToDomainQuestion())
		}

		return nil
	})

	return questions, err
}

func (r *QuestionRepository) Create(ctx context.Context, question *repositories.Question) error {
	dbQuestion, err := ToDatabaseQuestion(question, false)
	if err != nil {
		return err
	}

	err = r.gorm.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		_, err := getByID[databaseQuiz](tx, dbQuestion.QuizID)
		if err != nil {
			return err
		}

		return tx.Create(dbQuestion).Error
	})

	question.ID = dbQuestion.ID.String()

	return err
}

func (r *QuestionRepository) Update(ctx context.Context, question *repositories.Question) error {
	dbQuestion, err := ToDatabaseQuestion(question, true)
	if err != nil {
		return err
	}

	return r.gorm.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		existing, err := getByID[databaseQuestion](tx, dbQuestion.ID)
		if err != nil {
			return err
		}

		existing.Title = dbQuestion.Title
		existing.OptionA = dbQuestion.OptionA
		existing.OptionB = dbQuestion.OptionB
		existing.OptionC = dbQuestion.OptionC
		existing.OptionD = dbQuestion.OptionD
		existing.Answer = dbQuestion.Answer

		return tx.Save(existing).Error
	})
}

func (r *QuestionRepository) Delete(ctx context.Context, id repositories.QuestionID) error {
	return r.gorm.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&databaseQuestion{}, toUUID(id)).Error; err != nil {
			return err
		}
		return nil
	})
}
