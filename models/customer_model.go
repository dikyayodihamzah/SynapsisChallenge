package models

import (
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required, email" gorm:"unique"`
	Password string `json:"password" binding:"required"`
	Orders []Order
}

func (Customer) TableName() string {
	return "customers"
}