package services

import (
	"context"
	"kvizo-api/internal/repositories"

	"github.com/google/uuid"
)

type QuestionService struct {
	repository repositories.QuestionRepository
}

func NewQuestionService(r repositories.QuestionRepository) *QuestionService {
	return &QuestionService{repository: r}
}

func (s *QuestionService) GetByID(ctx context.Context, id string) (*repositories.Question, error) {
	questionUUID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	quiz, err := s.repository.GetByID(
		ctx,
		repositories.QuestionID(questionUUID))

	if err != nil {
		return nil, err
	}

	return quiz, nil
}

func (s *QuestionService) ListByQuizID(ctx context.Context, quizID string) ([]*repositories.Question, error) {
	quizUUID, err := uuid.Parse(quizID)
	if err != nil {
		return nil, err
	}

	return s.repository.ListByQuizID(
		ctx,
		repositories.QuizID(quizUUID))
}

func (s *QuestionService) Create(ctx context.Context, question *repositories.Question) error {
	return s.repository.Create(ctx, question)
}

func (s *QuestionService) Delete(ctx context.Context, id string) error {
	questionUUID, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	return s.repository.Delete(ctx, repositories.QuestionID(questionUUID))
}
