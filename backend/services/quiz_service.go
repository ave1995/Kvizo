package services

import (
	"context"
	"kvizo-api/internal/repositories"

	"github.com/google/uuid"
)

type QuizService struct {
	repository repositories.QuizRepository
}

// TODO: dělám tohle dobře? Zde posílám kopii. Je nějáký způsob jak to udělat přes referenci, když je to interface?
func NewQuizService(r repositories.QuizRepository) *QuizService {
	return &QuizService{repository: r}
}

// TODO: jak správně kontrolovat i v té horní metodě, něják moc ifů
func (s *QuizService) GetQuiz(ctx context.Context, id string) (*repositories.Quiz, error) {
	quizUUID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	quizID := repositories.QuizID(quizUUID)

	quiz, err := s.repository.GetByID(ctx, quizID)
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
