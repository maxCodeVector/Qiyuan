package main

import (
	"os"
	"qiyuan/service"
	"testing"
)

func TestUpload(t *testing.T) {

	beforeTest("Bob", "10086")
	defer afterTest("Bob", "10086")

	tempFileURL:= "/home/qydev/go/src/qiyuan/testgo.go"
	file, _ := os.Open(tempFileURL)
	newFileURL := service.UploadFile("10086", file, "hello")

	orderFromDB, _ := service.GetOrderByID("10086")

	if orderFromDB.FileUrl != newFileURL {
		t.Fail()
	}

}

func TestDownload(t *testing.T) {
	beforeTest("Bob", "10086")
	defer afterTest("Bob", "10086")

}

