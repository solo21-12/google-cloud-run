package models

import "github.com/google/uuid"

type Group struct {
    ID    int       `gorm:"primaryKey;autoIncrement" json:"id"`
    UID   uuid.UUID `gorm:"unique" json:"uid"` 
    Name  string    `json:"name"`
    Users []User    `gorm:"many2many:Groups_Users_Maps;foreignKey:ID;joinForeignKey:Groups_Id;References:ID;joinReferences:Users_Id" json:"users"`
}