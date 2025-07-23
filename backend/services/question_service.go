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

// TODO: jak správně kontrolovat i v té horní metodě, něják moc ifů
func (s *QuestionService) GetByID(ctx context.Context, id string) (*repositories.Question, error) {
	questionUUID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	parsedQuestionID := repositories.QuestionID(questionUUID)

	quiz, err := s.repository.GetByID(ctx, parsedQuestionID)
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

	parsedQuizID := repositories.QuizID(quizUUID)

	return s.repository.ListByQuizID(ctx, parsedQuizID)
}

func (s *QuestionService) Create(ctx context.Context, question *repositories.Question) error {
	return s.repository.Create(ctx, question)
}
