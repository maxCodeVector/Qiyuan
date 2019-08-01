package model

import (
	"db"
	"errors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
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
	conn := db.GetConnFromDB()

	var orders []Order
	// Get all records
	conn.Find(&orders)
	return &orders
}

func GetOrderByID(orderId string) (*Order, error) {
	conn := db.GetConnFromDB()

	order := new(Order)
	// Get first matched record
	err := conn.Where("order_id = ?", orderId).First(order).Error
	if err == nil{
		return order, nil
	}

	return nil, errors.New("Order not found")
}