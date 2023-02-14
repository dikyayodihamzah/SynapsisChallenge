package models

import "gorm.io/gorm"

type Category struct{
	gorm.Model
	Name string `json:"name" binding:"required"`
	Desc string `json:"description"`
}

func (Category) TableName() string {
	return "categories"
}