package infrastructure

import (
	"fmt"
	"strings"

	interfaces "github.com/google-run-code/Domain/Interfaces"
	models "github.com/google-run-code/Domain/Models"
	"github.com/google-run-code/config"
)

type JwtService struct {
	Env *config.Env
}

func NewJwtService(env *config.Env) interfaces.JwtService {
	return &JwtService{
		Env: env,
	}
}

func (j *JwtService) ValidateToken(tokenStr string) (*models.JWTCustome, error) {
	return nil, nil
}

func (j *JwtService) ValidateAuthHeader(authHeader string) ([]string, error) {
	if authHeader == "" {
		return nil, fmt.Errorf("authorization header is required")
	}

	authParts := strings.Split(authHeader, " ")
	if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" {
		return nil, fmt.Errorf("invalid authorization header")
	}

	return authParts, nil
}
