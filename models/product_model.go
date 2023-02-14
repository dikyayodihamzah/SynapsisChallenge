package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name  string `json:"name"`
	Desc  string `json:"desc"`
	Price int	 `json:"price"`
	CategoryID uint `json:"category"`
	Category Category
}

func (Product) TableName() string {
	return "products"
}