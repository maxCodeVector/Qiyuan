package main

import (
	"qiyuan/model"
	"qiyuan/service"
	"testing"
)


func TestCreateOrder1(t *testing.T)  {

	order := model.Order{OrderId:"1234", UserName:"xyz"}
	service.CreateOrder(&order)

	oneOrder, err := service.GetOrderByID("1234")
	if err != nil || oneOrder == nil {
		t.Fail()
	}

	service.DeleteOrder(oneOrder)
}

func TestGetOrderByID1(t *testing.T)  {

	oneOrder, err := service.GetOrderByID("2333")
	if err != nil || oneOrder == nil {
		t.Fail()
	}
}

func TestGetOrderByID2(t *testing.T)  {
	_, err := service.GetOrderByID("123")
	if err == nil{
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
	orderList := service.FuzzySearchOrder("hya", true, true)

	if orderList != nil && len(*orderList) >= 0 {
		return
	}
	t.Error("orderList must be iterable")
}


func TestDeleteOrder(t *testing.T) {
	// should not delete all record when order does not have primary key
	if service.DeleteOrder(&model.Order{}){
		t.Fail()
	}
	tempOrder := model.Order{ID:10086, UserName:"hello"}
	service.CreateOrder(&tempOrder)
	if !service.DeleteOrder(&tempOrder){
		t.Fail()
	}

}


func TestUpdateOrder(t *testing.T) {

	tempOrder := model.Order{OrderId:"10086", UserName:"hello", FileUrl:"www.youtube.com"}
	service.CreateOrder(&tempOrder)

	oneOrder, err := service.GetOrderByID("10086")
	if err != nil || oneOrder == nil || oneOrder.FileUrl != "www.youtube.com"{
		t.Fail()
	}
	oneOrder.FileUrl = "www.facebook.com"

	service.UpdateOrder(oneOrder)

	oneOrder, err = service.GetOrderByID("10086")
	if err != nil || oneOrder == nil || oneOrder.FileUrl != "www.facebook.com"{
		service.DeleteOrder(&tempOrder)
		t.Fail()
		return
	}
	service.DeleteOrder(&tempOrder)

}