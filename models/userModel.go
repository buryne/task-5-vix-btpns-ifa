package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string  `grom:"unique" grom:"not null"  binding:"required"`
	Email    string  `grom:"unique" grom:"not null" binding:"required"`
	Password string  `gorm:"not null"`
	Photo    []Photo `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
