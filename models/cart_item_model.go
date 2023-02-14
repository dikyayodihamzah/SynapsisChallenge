package models

import "gorm.io/gorm"

type CartItem struct {
	gorm.Model
	ProductID uint `gorm:"product_id" json:"product_id"`
	Product Product
	Quantity  int `json:"quantity"`
}

func (CartItem) TableName() string {
	return "cart_items"
}