package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWTManager interface {
	Generate(user *User) (string, error)
	Verify(tokenString string) (*User, error)
}

type jwtManager struct {
	secretKey     string
	tokenDuration time.Duration
}

func NewJWTManager(secretKey string, duration time.Duration) JWTManager {
	return &jwtManager{secretKey: secretKey, tokenDuration: duration}
}

type Claims struct {
	UserID string `json:"sub"`
	Email  string `json:"email"`
	// Role   string `json:"role"`
	jwt.RegisteredClaims
}

func (j *jwtManager) Generate(user *User) (string, error) {
	claims := Claims{
		UserID: user.ID.String(),
		Email:  user.Email,
		// Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.tokenDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (j *jwtManager) Verify(tokenString string) (*User, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Validate signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	parsed, err := uuid.Parse(claims.UserID)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID in claims: %q: %w", claims.UserID, err)
	}

	user := &User{
		ID:    parsed,
		Email: claims.Email,
		// Role:  claims.Role,
	}

	return user, nil
}
