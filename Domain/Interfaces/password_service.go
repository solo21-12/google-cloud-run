package interfaces

import models "github.com/google-run-code/Domain/Models"

type PasswordService interface {
	EncryptPassword(password string) (string, error)
	ValidatePassword(password string, hashedPassword string) bool
	ValidatePasswordStrength(password string) *models.ErrorResponse
}
