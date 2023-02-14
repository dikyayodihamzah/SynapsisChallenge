package models

import "gorm.io/gorm"

type CartItem struct {
	gorm.Model
	OrderID uint `json:"order_id"`
	ProductID uint `json:"product_id"`
	Product Product
	Quantity  int `json:"quantity"`
}

func (CartItem) TableName() string {
	return "cart_items"
}