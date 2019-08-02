package service

import (
	"errors"
	"qiyuan/db"
	"qiyuan/model"
)

func CreateOrder(order *model.Order){
	conn := db.GetConnFromDB("../test.sqlite")
	conn.Create(order)
}

func UpdateOrder(order *model.Order) bool{
	conn := db.GetConnFromDB("../test.sqlite")
	orderRecord, err := GetOrderByID(order.OrderId)
	if err != nil{
		return false
	}
	orderRecord.FileUrl = order.FileUrl
	orderRecord.Status = order.Status
	orderRecord.Amount = order.Amount
	conn.Model(orderRecord).Updates(orderRecord)
	return true
}


func DeleteOrder(order *model.Order) bool{
	conn := db.GetConnFromDB("../test.sqlite")
	if order.ID != 0 {
		conn.Delete(order)
		return true
	}
	return false
}

func FuzzySearchOrder(name string, orderByTime bool, orderByAmount bool)  *[]model.Order  {


	conn := db.GetConnFromDB("../test.sqlite")
	var orders []model.Order

	conn.Where("user_name = ?", name).Find(&orders)

	return &orders
}

func GetOrderByID(orderId string) (*model.Order, error) {
	conn := db.GetConnFromDB("../test.sqlite")

	order := new(model.Order)
	// Get first matched record
	err := conn.Where("order_id = ?", orderId).First(order).Error
	if err == nil {
		return order, nil
	}

	return nil, errors.New("Order not found")
}