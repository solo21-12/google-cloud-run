package usecases

import (
	interfaces "github.com/google-run-code/Domain/Interfaces"
	models "github.com/google-run-code/Domain/Models"
)

type generateTokenUsecase struct {
	jwtService interfaces.JwtService
}

func NewGenerateTokenUseCase(jwtservice interfaces.JwtService) interfaces.GenerateTokenUseCase {
	return generateTokenUsecase{
		jwtService: jwtservice,
	}
}

func (g generateTokenUsecase) GenerateAccessToken(database string) (*string, *models.ErrorResponse) {
	token, err := g.jwtService.GenerateToken(database)
	if err != nil {
		return nil, models.InternalServerError("Error generating token")
	}
	return &token, nil
}
