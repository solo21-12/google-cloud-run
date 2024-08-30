package infrastructure

import (
	"github.com/asaskevich/govalidator"
	interfaces "github.com/google-run-code/Domain/Interfaces"
	"github.com/google-run-code/config"
)

type emailService struct {
	env config.Env
}

func NewEmailService(env config.Env) interfaces.EmailService {
	return &emailService{
		env: env,
	}
}

func (es *emailService) IsValidEmail(email string) bool {
	return govalidator.IsEmail(email)
}
