package models

type Group struct {
	UID   uint   `gorm:"primaryKey" json:"uid"`
	Name  string `json:"name"`
	Users []User `gorm:"many2many:user_groups;" json:"users"`
}
