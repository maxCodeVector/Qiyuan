package main

import (
	"qiyuan/model"
	"qiyuan/service"
	"testing"
)

func TestGetOneOrder1(t *testing.T)  {

	oneOrder, err := service.GetOrderByID("2333")
	if err != nil || oneOrder == nil {
		t.Fail()
	}
}

func TestGetOneOrder2(t *testing.T)  {
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