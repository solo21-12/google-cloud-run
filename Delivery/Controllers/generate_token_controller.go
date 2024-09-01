package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	interfaces "github.com/google-run-code/Domain/Interfaces"
)

type generateAccessToken struct {
	generateTokenUseCase interfaces.GenerateTokenUseCase
}

type Database struct {
	Database string `json:"database"`
}

func NewGenerateTokenController(gat interfaces.GenerateTokenUseCase) interfaces.GenerateTokenController {
	return generateAccessToken{
		generateTokenUseCase: gat,
	}
}

func (g generateAccessToken) GenerateAccessToken(x *gin.Context) {
	var database Database

	if err := x.ShouldBindJSON(&database); err != nil {
		x.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := g.generateTokenUseCase.GenerateAccessToken(database.Database)
	if err != nil {
		x.JSON(err.Code, gin.H{"error": err.Error()})
		return
	}

	x.JSON(http.StatusOK, gin.H{"token": token})
}
