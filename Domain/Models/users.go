package models

import "github.com/google/uuid"

type User struct {
	UID      uuid.UUID `gorm:"primaryKey" json:"uid"`
	Name     string    `json:"name"`
	Email    string    `gorm:"unique" json:"email"`
	Status   int       `json:"status"`
	Groups   []*Group  `gorm:"many2many:user_groups;" json:"groups"`
	Roles    []*Role   `gorm:"foreignKey:UserID" json:"roles"`
}
