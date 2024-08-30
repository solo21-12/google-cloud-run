package models

import (
	"github.com/google/uuid"
)

type Session struct {
	ID           uuid.UUID `gorm:"type:uuid;" json:"id"`
	UserID       string    `gorm:"type:uuid;not null;" json:"user_id"`
	RefreshToken string    `gorm:"type:text;not null;" json:"refresh_token"`
	AccessToken  string    `gorm:"type:text;not null;" json:"access_token"`
}
