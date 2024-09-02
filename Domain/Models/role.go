package models

import (
	"encoding/json"

	"github.com/google/uuid"
)

type Role struct {
	ID     int             `gorm:"primaryKey;autoIncrement" json:"id"`
	UID    uuid.UUID       `gorm:"unique" json:"uid"` 
	Name   string          `json:"name"`
	Rights json.RawMessage `gorm:"type:json" json:"rights"`
	Users  []*User         `gorm:"foreignKey:RoleID" json:"users"` 
}
