package services

import (
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
func (s *QuizService) GetByID(id string) (*repositories.Quiz, error) {
	quizID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	quiz, err := s.repository.GetByID(quizID)
	if err != nil {
		return nil, err
	}

	return quiz, nil
}

func (s *QuizService) List() ([]*repositories.Quiz, error) {
	return s.repository.List()
}

func (s *QuizService) Create(quiz *repositories.Quiz) error {
	return s.repository.Create(quiz)
}
