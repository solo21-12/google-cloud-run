package models

import "github.com/google/uuid"

type Group struct {
	GID   uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"gid"`
	Name  string    `json:"name"`
	Users []User    `gorm:"many2many:user_groups;" json:"users"`
}
