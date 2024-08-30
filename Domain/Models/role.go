package models

import "github.com/google/uuid"

type Rights struct {
	Create bool `json:"create"`
	Read   bool `json:"read"`
	Update bool `json:"update"`
	Delete bool `json:"delete"`
}

type Role struct {
	UID    uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"uid"`
	Name   string    `json:"name"`
	Rights Rights    `gorm:"type:json" json:"rights"`
	UserID uuid.UUID `json:"user_id"`
	User   User      `gorm:"foreignKey:UserID" json:"user"`
}
