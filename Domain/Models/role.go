package models

type Role struct {
	UID    uint   `gorm:"primaryKey" json:"uid"`
	Name   string `json:"name"`
	Rights string `gorm:"type:json" json:"rights"`
	UserID uint   `json:"user_id"`
	User   User   `gorm:"foreignKey:UserID" json:"user"`
}
