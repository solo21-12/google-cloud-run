package infrastructure_test

import (
	"testing"

	interfaces "github.com/google-run-code/Domain/Interfaces"
	infrastructure "github.com/google-run-code/Infrastructure"
	"github.com/google-run-code/config"
	"github.com/stretchr/testify/suite"
)

type EmailServiceTestSuite struct {
	suite.Suite
	emailService interfaces.EmailService
}

func (suite *EmailServiceTestSuite) SetupTest() {
	env := config.Env{}
	suite.emailService = infrastructure.NewEmailService(env)
}

func (suite *EmailServiceTestSuite) TestIsValidEmail_ValidEmail() {
	validEmail := "test@example.com"
	isValid := suite.emailService.IsValidEmail(validEmail)
	suite.True(isValid, "Expected email to be valid")
}

func (suite *EmailServiceTestSuite) TestIsValidEmail_InvalidEmail() {
	invalidEmail := "invalid-email"
	isValid := suite.emailService.IsValidEmail(invalidEmail)
	suite.False(isValid, "Expected email to be invalid")
}

func TestEmailServiceTestSuite(t *testing.T) {
	suite.Run(t, new(EmailServiceTestSuite))
}
