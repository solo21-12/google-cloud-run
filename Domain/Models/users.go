package models

import "github.com/google/uuid"

type User struct {
	ID     int        `gorm:"primaryKey;autoIncrement" json:"id"`
	UID    uuid.UUID  `gorm:"unique" json:"uid"`
	Name   string     `json:"name"`
	Email  string     `gorm:"unique" json:"email"`
	Status int        `json:"status"`
	Groups []*Group   `gorm:"many2many:Groups_Users_Map;" json:"groups"`
	RoleID *int       `gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"role_id"`
	Role   *Role      `gorm:"foreignKey:RoleID;references:ID" json:"role,omitempty"`
}

