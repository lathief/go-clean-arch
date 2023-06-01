package entity

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Name     string `gorm:"column:name"`
	Email    string `gorm:"column:email"`
	Password string `gorm:"column:password"`
}
