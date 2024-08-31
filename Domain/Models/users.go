package models

import "github.com/google/uuid"

type User struct {
	UID    uuid.UUID  `gorm:"primaryKey" json:"uid"`
	Name   string     `json:"name"`
	Email  string     `gorm:"unique" json:"email"`
	Status int        `json:"status"`
	Groups []*Group   `gorm:"many2many:user_groups;" json:"groups"`
	RoleID *uuid.UUID `gorm:"index" json:"role_id"` // RoleID as a pointer to allow NULL
	Role   *Role      `gorm:"foreignKey:RoleID" json:"role,omitempty"` // Role is optional
}