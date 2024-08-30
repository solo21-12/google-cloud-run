package models

import "github.com/google/uuid"

type Group struct {
	GID   uuid.UUID `gorm:"primaryKey" json:"gid"`
	Name  string    `json:"name"`
	Users []User    `gorm:"many2many:user_groups;" json:"users"`
}
