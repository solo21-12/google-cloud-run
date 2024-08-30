package interfaces

import models "github.com/google-run-code/Domain/Models"

type JwtService interface {
	CreateAccessToken(user models.User, expTime int) (accessToken string, err error)
	CreateRefreshToken(user models.User, expTime int) (refreshToken string, err error)
	ValidateToken(tokenStr string) (*models.JWTCustome, error)
	ValidateAuthHeader(authHeader string) ([]string, error)
}
