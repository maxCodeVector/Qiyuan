package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"qiyuan/db"
)

type Order struct {
	gorm.Model
	OrderId  string
	UserName string
	Amount   float64
	Status   string
	FileUrl  string
}

// set User's table name to be `profiles`
func (Order) TableName() string {
	return "demo_order"
}

func GetOrders() *[]Order {
	conn := db.GetConnFromDB("../test.sqlite")

	var orders []Order
	// Get all records
	conn.Find(&orders)
	return &orders
}