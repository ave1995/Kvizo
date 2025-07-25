package auth

import (
	"context"
)

type AuthService struct {
	repository AuthRepository
	jwtManager JWTManager
}

func NewAuthService(r AuthRepository, j JWTManager) *AuthService {
	return &AuthService{repository: r, jwtManager: j}
}

func (s *AuthService) AuthenticateUser(ctx context.Context, email, password string) (string, error) {
	user, err := s.repository.AuthenticateUser(ctx, email, password)
	if err != nil {
		return "", err
	}

	return s.jwtManager.Generate(user)
}

func (s *AuthService) RegisterUser(ctx context.Context, email, password string) (*User, error) {
	return s.repository.RegisterUser(ctx, email, password)
}
