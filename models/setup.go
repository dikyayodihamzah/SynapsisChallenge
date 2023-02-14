package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB


func ConnectDatabase() {
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/synapsis_challenge?parseTime=true"))
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(
		&Customer{}, 
		&Category{}, 
		&Product{}, 
		&CartItem{},
		&Order{},
	)
	DB = db
}