package auth

import "github.com/gin-gonic/gin"

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
