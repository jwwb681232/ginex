package models

import (
	"github.com/jinzhu/gorm"
)

type HomeCook struct {
	gorm.Model
	Name          string  `gorm:"column:name"`
	Email         string  `gorm:"column:email"`
	Password      string  `gorm:"column:password"`
	RememberToken *string `gorm:"column:remember_token"`
}