package models

import "github.com/google/uuid"

type Rights struct {
	Create bool `json:"create"`
	Read   bool `json:"read"`
	Update bool `json:"update"`
	Delete bool `json:"delete"`
}

type Role struct {
	RID    uuid.UUID `gorm:"primaryKey" json:"gid"`
	Name   string    `json:"name"`
	Rights Rights    `gorm:"type:json" json:"rights"`
	UserID uuid.UUID `json:"user_id"`
	User   User      `gorm:"foreignKey:UserID" json:"user"`
}
