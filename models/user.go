package models

import (
	"ginex/database"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name          string  `gorm:"column:name"`
	Email         string  `gorm:"column:email"`
	Password      string  `gorm:"column:password"`
	RememberToken *string `gorm:"column:remember_token"`
}

func (User) TableName() string {
	return "users"
}

func (User) WhereEmail(email *string) (User, bool) {
	var userData User
	var notFound bool

	notFound = database.Db.Where("email = ?", email).First(&userData).RecordNotFound()
	return userData, notFound
}

func (User) CreateUser(u User) *gorm.DB {
	return database.Db.Create(&u)
}
