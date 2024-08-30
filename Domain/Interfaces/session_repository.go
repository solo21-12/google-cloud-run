package interfaces

import (
	"context"

	models "github.com/google-run-code/Domain/Models"
)

type SessionRepository interface {
	SaveToken(ctx context.Context, session *models.Session) *models.ErrorResponse
	UpdateToken(ctx context.Context, session *models.Session) *models.ErrorResponse
	RemoveToken(ctx context.Context, userID string) *models.ErrorResponse
	GetToken(ctx context.Context, userID string) (*models.Session, *models.ErrorResponse)
}
