package model

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	Name     string `gorm:"column:name"`
	Email    string `gorm:"column:email"`
	Password string `gorm:"column:password"`
}
