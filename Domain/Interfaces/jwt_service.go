package interfaces

import models "github.com/google-run-code/Domain/Models"

type JwtService interface {
	ValidateToken(tokenStr string) (*models.JWTCustome, error)
	ValidateAuthHeader(authHeader string) ([]string, error)
	GenerateToken(database string) (string, error)
}
