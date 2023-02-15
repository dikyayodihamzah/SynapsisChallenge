package models

import "gorm.io/gorm"

type CartItem struct {
	gorm.Model
	CustomerID uint `json:"customer_id" binding:"required"`
	Customer   Customer
	ProductID  uint `json:"product_id" binding:"required"`
	Product    Product
	OrderID    uint `json:"order_id" binding:"required"`
	Quantity   int  `json:"quantity" binding:"required"`
}

func (CartItem) TableName() string {
	return "cart_items"
}
