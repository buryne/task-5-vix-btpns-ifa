package models

import "gorm.io/gorm"


type Photo struct {
	gorm.Model
	Title    string `gorm:"not null"`
	Caption  string `gorm:"not null"`
	PhotoUrl string `gorm:"not null"`
	UserID   int   `gorm:"not null"`
	User     User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
