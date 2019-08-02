package service

import (
	"io"
	"os"
	"qiyuan/model"
)

func UploadFile(orderId string, file io.Reader, fileName string) string {

	out, err := os.Create("../file/" + fileName)
	if err != nil {
		return ""
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		return ""
	}
	filePath := "http://127.0.0.1:8000/file/" + fileName

	if order, err := GetOrderByID(orderId); err == nil{
		order.FileUrl = filePath
		UpdateOrder(order)
		return filePath
	}
	return ""
}

func DownloadFile(order *model.Order) bool{
	return false
}