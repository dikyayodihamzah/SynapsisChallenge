package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	CustomerID uint `json:"customer_id"`
	Customer Customer
	CartItems []CartItem
	Total	   uint `json:"total"`
	Paid       bool `json:"paid"`
}

func (Order) TableName() string {
	return "orders"
}