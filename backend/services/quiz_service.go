package services

import (
	"context"
	"kvizo-api/internal/repositories"

	"github.com/google/uuid"
)

type QuizService struct {
	repository repositories.QuizRepository
}

func NewQuizService(r repositories.QuizRepository) *QuizService {
	return &QuizService{repository: r}
}

func (s *QuizService) GetQuiz(ctx context.Context, id string) (*repositories.Quiz, error) {
	quizUUID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	quiz, err := s.repository.GetByID(ctx, repositories.QuizID(quizUUID))
	if err != nil {
		return nil, err
	}

	return quiz, nil
}

func (s *QuizService) List(ctx context.Context) ([]*repositories.Quiz, error) {
	return s.repository.List(ctx)
}

func (s *QuizService) Create(ctx context.Context, quiz *repositories.Quiz) error {
	return s.repository.Create(ctx, quiz)
}

func (s *QuizService) Delete(ctx context.Context, id string) error {
	quizUUID, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	return s.repository.Delete(ctx, repositories.QuizID(quizUUID))
}
