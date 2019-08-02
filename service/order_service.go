package service

import (
	"errors"
	"qiyuan/db"
	"qiyuan/model"
)

func CreateOrder(){

}

func UpdateOrder(){

}

func QueryOrderByTime()  {
	
}

func QueryOrderByAmount()  {

}

func FuzzySearchOrder()  {

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