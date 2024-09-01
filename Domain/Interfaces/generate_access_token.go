package interfaces

import (
	"github.com/gin-gonic/gin"
	models "github.com/google-run-code/Domain/Models"
)

type GenerateTokenUseCase interface {
	GenerateAccessToken(database string) (*string, *models.ErrorResponse)
}

type GenerateTokenController interface {
	GenerateAccessToken(x *gin.Context)
}
