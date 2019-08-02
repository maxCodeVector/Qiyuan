package main

import (
	"qiyuan/db"
	"qiyuan/model"
	"qiyuan/service"
	"strconv"
	"testing"
)

func beforeTest(tempName, tempOrderId string) {
	conn := db.GetConnFromDB("../test.sqlite")
	order := model.Order{OrderId: tempOrderId, UserName: "xyz", FileUrl:"www.youtube.com"}
	conn.Create(&order)
	var orders [5]model.Order
	for i, order := range orders {
		order.UserName = tempName
		order.Status = "NK"
		order.OrderId = strconv.Itoa(10087 + i)
		order.Amount = 1000
		conn.Save(&order)
	}
}

func afterTest(tempName, tempOrderId string) {
	conn := db.GetConnFromDB("../test.sqlite")
	conn.Exec("delete from demo_order where user_name = ?", tempName)
	conn.Exec("delete from demo_order where order_id = ?", tempOrderId)
}

func TestCreateOrder1(t *testing.T) {

	order := model.Order{OrderId: "10086", UserName: "xyz"}
	service.CreateOrder(&order)

	oneOrder, err := service.GetOrderByID("10086")
	if err != nil || oneOrder == nil {
		t.Fail()
	}

	service.DeleteOrder(oneOrder)
}

func TestGetOrderByID1(t *testing.T) {

	beforeTest("Bob", "10086")
	defer afterTest("Bob", "10086")

	oneOrder, err := service.GetOrderByID("10086")
	if err != nil || oneOrder == nil {
		t.Fail()
	}
}

func TestGetOrderByID2(t *testing.T) {
	_, err := service.GetOrderByID("10086")
	if err == nil {
		t.Fail()
	}
}

// Test the function that fetches all articles
func TestGetAllOrders(t *testing.T) {
	orderList := model.GetOrders()

	if orderList != nil && len(*orderList) >= 0 {
		return
	}
	t.Error("orderList must be iterable")
}

func TestFuzeQuery(t *testing.T) {

	beforeTest("Bob", "10086")
	defer afterTest("Bob", "10086")

	orderList := service.FuzzySearchOrder("Bob", true, true)

	if orderList != nil && len(*orderList) >= 0 {
		return
	}
	t.Error("orderList must be iterable")
}

func TestDeleteOrder(t *testing.T) {

	beforeTest("Bob", "10086")
	defer afterTest("Bob", "10086")

	// should not delete all record when order does not have primary key
	if service.DeleteOrder(&model.Order{}) {
		t.Fail()
	}

	spOrder, _ := service.GetOrderByID("10086")
	if !service.DeleteOrder(spOrder) {
		t.Fail()
	}
	if _, err2 := service.GetOrderByID("10086"); err2 == nil{
		t.Fail()
	}

}

func TestUpdateOrder(t *testing.T) {

	beforeTest("Bob", "10086")
	defer afterTest("Bob", "10086")

	oneOrder, err := service.GetOrderByID("10086")
	if err != nil || oneOrder == nil || oneOrder.FileUrl != "www.youtube.com" {
		t.Fail()
	}
	oneOrder.FileUrl = "www.facebook.com"

	service.UpdateOrder(oneOrder)

	oneOrder, err = service.GetOrderByID("10086")
	if err != nil || oneOrder == nil || oneOrder.FileUrl != "www.facebook.com" {
		t.Fail()
	}
}

func TestCheckOut(t *testing.T) {

	beforeTest("Bob", "10086")
	defer afterTest("Bob", "10086")

	totalAmount := service.CheckOut("Bob")
	if totalAmount != 5000 {
		t.Fail()
		return
	}

	totalAmount = service.CheckOut("Bob")
	if totalAmount != 0 {
		t.Fail()
	}

}

