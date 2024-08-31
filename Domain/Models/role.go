package models

import (
	"encoding/json"

	"github.com/google/uuid"
)

type Role struct {
	RID uuid.UUID `gorm:"primaryKey" json:"rid"`
	Name   string          `json:"name"`
	Rights json.RawMessage `gorm:"type:json" json:"rights"`
	Users  []*User         `gorm:"foreignKey:RoleID" json:"users"` // Role may have many users
}
