package models

import "gorm.io/gorm"

type CartItem struct {
	gorm.Model
	OrderID uint `json:"order_id" binding:"required"`
	ProductID uint `json:"product_id" binding:"required"`
	Product Product
	Quantity  int `json:"quantity" binding:"required"`
}

func (CartItem) TableName() string {
	return "cart_items"
}