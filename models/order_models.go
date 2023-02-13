package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	Id         int64 `gorm: "primaryKey" json:"id"`
	CustomerId int64 `gorm: "foreignKey" json:"id"`
	CartItemId int64 `gorm: "foreignKey" json:"id"`
	Paid       bool  `gorm: "default:false" json: "paid"`
}