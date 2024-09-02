package models

import "github.com/google/uuid"

type Group struct {
    ID    int       `gorm:"primaryKey;autoIncrement" json:"id"`
    UID   uuid.UUID `gorm:"unique" json:"uid"` 
    Name  string    `json:"name"`
    Users []User    `gorm:"many2many:Groups_Users_Maps;" json:"users"`
}