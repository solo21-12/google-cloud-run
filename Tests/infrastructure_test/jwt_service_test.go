package infrastructure_test

import (
	"strings"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	interfaces "github.com/google-run-code/Domain/Interfaces"
	models "github.com/google-run-code/Domain/Models"
	infrastructure "github.com/google-run-code/Infrastructure"
	"github.com/google-run-code/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type JwtServiceTestSuite struct {
	suite.Suite
	service interfaces.JwtService
	env     *config.Env
}

func (suite *JwtServiceTestSuite) SetupTest() {
	suite.env = &config.Env{
		JWT_SECRET: "sample-token-secret",
	}
	suite.service = infrastructure.NewJwtService(suite.env)
}

func (suite *JwtServiceTestSuite) TestValidateToken_Success() {
	tokenString, err := suite.service.GenerateToken("test-database")
	assert.NoError(suite.T(), err)

	claims, err := suite.service.ValidateToken(tokenString)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "test-database", claims.Database)
}

func (suite *JwtServiceTestSuite) TestValidateToken_Expired() {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &models.JWTCustome{
		Expires:  time.Now().Add(-time.Hour).Unix(),
		Database: "test-database",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(-time.Hour).Unix(),
		},
	})
	tokenString, err := token.SignedString([]byte(suite.env.JWT_SECRET))
	assert.NoError(suite.T(), err)

	claims, err := suite.service.ValidateToken(tokenString)
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), claims)
}

func (suite *JwtServiceTestSuite) TestValidateAuthHeader_Success() {
	authHeader := "Bearer some-token"
	authParts, err := suite.service.ValidateAuthHeader(authHeader)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "Bearer", authParts[0])
	assert.Equal(suite.T(), "some-token", authParts[1])
}

func (suite *JwtServiceTestSuite) TestValidateAuthHeader_InvalidFormat() {
	authHeader := "InvalidHeader"
	authParts, err := suite.service.ValidateAuthHeader(authHeader)
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), authParts)
	assert.Equal(suite.T(), "invalid authorization header", err.Error())
}

func (suite *JwtServiceTestSuite) TestValidateAuthHeader_Empty() {
	authHeader := ""
	authParts, err := suite.service.ValidateAuthHeader(authHeader)
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), authParts)
	assert.Equal(suite.T(), "authorization header is required", err.Error())
}

func (suite *JwtServiceTestSuite) TestGenerateToken_Success() {
	tokenString, err := suite.service.GenerateToken("test-database")
	assert.NoError(suite.T(), err)
	assert.True(suite.T(), strings.HasPrefix(tokenString, "eyJ"))
}

func TestJwtServiceTestSuite(t *testing.T) {
	suite.Run(t, new(JwtServiceTestSuite))
}
