package services

import (
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
func (s *QuestionService) GetByID(id string) (*repositories.Question, error) {
	parsedQuestionID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	quiz, err := s.repository.GetByID(parsedQuestionID)
	if err != nil {
		return nil, err
	}

	return quiz, nil
}

func (s *QuestionService) ListByQuizID(quizID string) ([]*repositories.Question, error) {
	parsedQuizID, err := uuid.Parse(quizID)
	if err != nil {
		return nil, err
	}

	return s.repository.ListByQuizID(parsedQuizID)
}

func (s *QuestionService) Create(question *repositories.Question) error {
	return s.repository.Create(question)
}
