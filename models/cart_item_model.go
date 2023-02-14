package models

import "gorm.io/gorm"

type CartItem struct {
	gorm.Model
	// Id        int64  `gorm: "primaryKey" json:"id"`
	ProductId int64 `gorm:"foreignKey" json:"product_id"`
	Quantity  int64  `json:"quantity"`
}