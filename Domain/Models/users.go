package models

import "github.com/google/uuid"

type User struct {
	UID    uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"uid"`
	Name   string    `json:"name"`
	Email  string    `gorm:"unique" json:"email"`
	Status int       `json:"status"`
	Password string  `json:"password"`
	Groups []*Group  `gorm:"many2many:user_groups;" json:"groups"`
	Roles  []*Role   `gorm:"foreignKey:UserID" json:"roles"`
}
