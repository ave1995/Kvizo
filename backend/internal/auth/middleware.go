package auth

import (
	"errors"

	"github.com/gin-gonic/gin"
)

const (
	UserContextKey = "user"
)

func AuthMiddleware(jwtManager JWTManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Authorization header missing"})
			return
		}

		const prefix = "Bearer "
		if len(authHeader) <= len(prefix) || authHeader[:len(prefix)] != prefix {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid authorization header format"})
			return
		}

		tokenString := authHeader[len(prefix):]

		user, err := jwtManager.Verify(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid or expired token"})
			return
		}

		// Inject into context for downstream handlers
		c.Set(UserContextKey, user)

		c.Next()
	}
}

var ErrUserNotInContext = errors.New("user not found in context")
var ErrUserHasInvalidTypeInContext = errors.New("user in context has invalid type")

func GetUserFromContext(c *gin.Context) (*User, error) {
	val, exists := c.Get(UserContextKey)
	if !exists {
		return nil, ErrUserNotInContext
	}

	user, ok := val.(*User)
	if !ok {
		return nil, ErrUserHasInvalidTypeInContext
	}

	return user, nil
}
