package routers

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/google-run-code/Delivery/Controllers"
	infrastructure "github.com/google-run-code/Infrastructure"
	usecases "github.com/google-run-code/Usecases"
	"github.com/google-run-code/config"
)

func NewGenerateTokenRouter(router *gin.RouterGroup) {
	env := config.NewEnv()
	jwtService := infrastructure.NewJwtService(env)

	generateTokenUseCase := usecases.NewGenerateTokenUseCase(jwtService)
	generateTokenController := controllers.NewGenerateTokenController(generateTokenUseCase)

	router.POST("/generate-token", generateTokenController.GenerateAccessToken)
}
