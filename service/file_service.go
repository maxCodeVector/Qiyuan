package service

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"io"
	"os"
	"qiyuan/model"
	"strconv"
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
	filePath := "file/" + fileName

	if order, err := GetOrderByID(orderId); err == nil{
		order.FileUrl = filePath
		UpdateOrder(order)
		return filePath
	}
	return ""
}

func ExportData(outFile string){

	orders := model.GetOrders()
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("sheet1")
	if err != nil {
		fmt.Printf(err.Error())
	}

	// add header
	row := sheet.AddRow()
	cell := row.AddCell()

	cell.Value = "OrderId"

	cell = row.AddCell()
	cell.Value = "UserName"

	cell = row.AddCell()
	cell.Value = "Amount"

	cell = row.AddCell()
	cell.Value = "Status"

	cell = row.AddCell()
	cell.Value = "FileUrl"

	//add data
	for _, order := range *orders{
		row = sheet.AddRow()
		cell = row.AddCell()
		cell.Value = order.OrderId

		cell = row.AddCell()
		cell.Value = order.UserName

		cell = row.AddCell()
		cell.Value = strconv.FormatFloat(order.Amount, 'f', 6, 64)

		cell = row.AddCell()
		cell.Value = order.Status

		cell = row.AddCell()
		cell.Value = order.FileUrl
	}
	err = file.Save(outFile)
	if err != nil {
		fmt.Printf(err.Error())
	}
	fmt.Println("\n\nexport success")

}