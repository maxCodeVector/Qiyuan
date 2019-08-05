package main

import (
	"os"
	"qiyuan/service"
	"testing"
)

func TestUpload(t *testing.T) {

	beforeTest("Bob", "10086")
	defer afterTest("Bob", "10086")

	tempFileURL:= "/home/qydev/go/src/qiyuan/file/testgo.go"
	file, _ := os.Open(tempFileURL)
	newFileURL := service.UploadFile("10086", file, "hello")

	orderFromDB, _ := service.GetOrderByID("10086")

	if orderFromDB.FileUrl != newFileURL {
		t.Fail()
	}

}

func TestExportData(t *testing.T) {

	beforeTest("Bob", "10086")
	defer afterTest("Bob", "10086")

	tempPath := "/tmp/export.xlsx"
	service.ExportData(tempPath)
	_, err := os.Stat(tempPath)
	if err != nil {
		t.Fail()
	}else {
		os.Remove(tempPath)
	}

}
