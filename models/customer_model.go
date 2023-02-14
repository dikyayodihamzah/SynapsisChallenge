package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	// Id       int64  `gorm: "primaryKey" json:"id"`
	Name     string `gorm:"type:varchar(300)" json:"name"`
	Email    string `gorm:"type:varchar(300)" json:"email"`
	Password string `gorm:"type:varchar(45)" json:"password"`
}