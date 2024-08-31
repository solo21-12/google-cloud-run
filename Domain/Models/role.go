package models

import (
	"encoding/json"

	"github.com/google/uuid"
)

type Role struct {
	ID     int             `gorm:"primaryKey;autoIncrement" json:"id"`
	RID    uuid.UUID       `gorm:"unique" json:"rid"` // Ensure RID is unique
	Name   string          `json:"name"`
	Rights json.RawMessage `gorm:"type:json" json:"rights"`
	Users  []*User         `gorm:"foreignKey:RoleID" json:"users"` // Role may have many users
}
