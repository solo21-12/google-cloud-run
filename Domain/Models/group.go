package models

import "github.com/google/uuid"

type Group struct {
    ID    int       `gorm:"primaryKey;autoIncrement" json:"id"`
    GID   uuid.UUID `gorm:"unique" json:"gid"` // Ensure GID is unique
    Name  string    `json:"name"`
    Users []User    `gorm:"many2many:Groups_Users_Map;" json:"users"`
}