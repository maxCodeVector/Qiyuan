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

	res := conn.Where("user_name like ?", "%"+name+"%")
	if orderByTime {
		res = res.Order("updated_at desc")
	}
	if orderByAmount {
		res = res.Order("amount desc")
	}
	res.Find(&orders)
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

/**
* check out all none OK order for this person with userName @userName, return the total checkout amount
*/
func CheckOut(userName string) float64 {
	conn := db.GetConnFromDB("../test.sqlite")
	tx := conn.Begin()

	var orders [] model.Order
	tx.Where("user_name = ? and status != ?", userName, "OK").Find(&orders)
	totalAmount := 0.

	for _, order := range orders{
		totalAmount += order.Amount
		order.Status = "OK"
		tx.Model(order).Updates(order)
	}

	tx.Commit()

	return totalAmount
}