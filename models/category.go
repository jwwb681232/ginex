package models

import "github.com/jinzhu/gorm"

type Category struct{
	gorm.Model
	Name	string	`gorm:"column:name"`
	Sort	uint	`gorm:"column:sort"`
}

func (Category) TableName() string {
	return "category"
}