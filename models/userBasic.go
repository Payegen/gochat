package models

import "gorm.io/gorm"

type UserBasic struct {
	gorm.Model
	Name     string
	PassWord string
	Phone    string
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}
