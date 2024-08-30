package models

type Rights struct {
	Create bool `json:"create"`
	Read   bool `json:"read"`
	Update bool `json:"update"`
	Delete bool `json:"delete"`
}

type Role struct {
	UID    uint   `gorm:"primaryKey" json:"uid"`
	Name   string `json:"name"`
	Rights Rights `gorm:"type:json" json:"rights"`
	UserID uint   `json:"user_id"`
	User   User   `gorm:"foreignKey:UserID" json:"user"`
}
