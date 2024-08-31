package middleware

import (
	"github.com/gin-gonic/gin"
	interfaces "github.com/google-run-code/Domain/Interfaces"
	"github.com/google-run-code/config"
)

func AuthMiddleware(jwtService interfaces.JwtService, env *config.Env) gin.HandlerFunc {
	return func(c *gin.Context) {
		// authHeader := c.GetHeader("Authorization")
		// authParts, err := jwtService.ValidateAuthHeader(authHeader)
		// if err != nil {
		// 	c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		// 	c.Abort()
		// 	return
		// }

		// tokenStr := authParts[1]
		// claims, err := jwtService.ValidateToken(tokenStr)
		// if err != nil {
		// 	c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		// 	c.Abort()
		// 	return
		// }

		// // Set database name in gin.Context
		// // if claims.DatabaseName == "" {
		// // 	c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid database name in token"})
		// // 	c.Abort()
		// // 	return
		// // }

		// // c.Set("dbName", claims.DatabaseName)
		// // c.Next()
	}
}
