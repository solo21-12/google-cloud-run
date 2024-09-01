package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	interfaces "github.com/google-run-code/Domain/Interfaces"
	"github.com/google-run-code/config"
)

func DatabaseMiddleware(env *config.Env, jwtService interfaces.JwtService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		authParts, err := jwtService.ValidateAuthHeader(authHeader)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		tokenStr := authParts[1]
		claims, err := jwtService.ValidateToken(tokenStr)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		// Get database name from token claims
		dbName := claims.Database
		if dbName == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Database name missing in token"})
			c.Abort()
			return
		}

		// Create a new database client for the requested database
		db := config.NewPostgresConfig(*env)
		client := db.Client(dbName)

		// Set the database client in the context
		c.Set("dbClient", client)

		c.Next()
	}
}
