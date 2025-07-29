package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type AuthService struct {
	repository           AuthRepository
	jwtManager           JWTManager
	refreshTokenDuration time.Duration
}

func NewAuthService(repository AuthRepository, jwtManager JWTManager, refreshTokenDuration time.Duration) *AuthService {
	return &AuthService{repository: repository, jwtManager: jwtManager, refreshTokenDuration: refreshTokenDuration}
}

func (s *AuthService) AuthenticateUser(ctx context.Context, email, password string) (*User, error) {
	return s.repository.AuthenticateUser(ctx, email, password)
}

func (s *AuthService) RegisterUser(ctx context.Context, email, password string) (*User, error) {
	return s.repository.RegisterUser(ctx, email, password)
}

func (s *AuthService) Login(ctx context.Context, email, password string) (*AccessToken, *RefreshToken, error) {
	user, err := s.repository.AuthenticateUser(ctx, email, password)
	if err != nil {
		return nil, nil, err
	}

	accessToken, err := s.jwtManager.Generate(user)
	if err != nil {
		return nil, nil, err
	}

	refreshToken := &RefreshToken{
		ID:        uuid.New(),
		UserID:    user.ID,
		Token:     uuid.New(), // random UUID token string
		ExpiresAt: time.Now().Add(s.refreshTokenDuration),
		Revoked:   false,
		CreatedAt: time.Now(),
	}

	err = s.repository.SaveRefreshToken(ctx, *refreshToken)
	if err != nil {
		return nil, nil, err
	}

	return accessToken, refreshToken, nil
}

func (s *AuthService) Refresh(ctx context.Context, refreshTokenString string) (*AccessToken, *RefreshToken, error) {
	// 1. Get refresh token from DB
	dbToken, err := s.repository.GetRefreshToken(ctx, refreshTokenString)
	if err != nil {
		return nil, nil, fmt.Errorf("refresh token not found or DB error: %w", err)
	}

	// 2. Check if token revoked or expired
	if dbToken.Revoked {
		return nil, nil, fmt.Errorf("refresh token revoked")
	}
	if time.Now().After(dbToken.ExpiresAt) {
		return nil, nil, fmt.Errorf("refresh token expired")
	}

	// 3. Get user associated with refresh token
	user, err := s.repository.GetUserByID(ctx, dbToken.UserID)
	if err != nil {
		return nil, nil, fmt.Errorf("user not found: %w", err)
	}

	// 4. Generate new access token
	accessToken, err := s.jwtManager.Generate(user)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to generate access token: %w", err)
	}

	// 5. Generate new refresh token
	newRefreshToken := RefreshToken{
		ID:        uuid.New(),
		UserID:    user.ID,
		Token:     uuid.New(),
		ExpiresAt: time.Now().Add(s.refreshTokenDuration),
		Revoked:   false,
		CreatedAt: time.Now(),
	}

	// 6. Save new refresh token to DB
	if err := s.repository.SaveRefreshToken(ctx, newRefreshToken); err != nil {
		return nil, nil, fmt.Errorf("failed to save new refresh token: %w", err)
	}

	// 7. Revoke old refresh token
	if err := s.repository.RevokeRefreshToken(ctx, refreshTokenString); err != nil {
		return nil, nil, fmt.Errorf("failed to revoke old refresh token: %w", err)
	}

	// 8. Return new tokens
	return accessToken, &newRefreshToken, nil
}
