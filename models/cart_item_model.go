package models

import "gorm.io/gorm"

type CartItem struct {
	gorm.Model
	Id        int64  `gorm: "primaryKey" json:"id"`
	ProductId string `gorm: "type:varchar(300)" json:"product_id"`
	Quantity  int64  `json: "quantity"`
}