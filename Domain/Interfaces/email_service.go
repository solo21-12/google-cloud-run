package interfaces

type EmailService interface {
	IsValidEmail(email string) bool
}
