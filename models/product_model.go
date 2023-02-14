package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model

	// Id    int64  `gorm:"primaryKey" json:"id"`
	Name  string `gorm:"type:varchar(300)" json:"name"`
	Desc  string `gorm:"type:varchar(3000)" json:"desc"`
	Price int64  `json:"price"`
}