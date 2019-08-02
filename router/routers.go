package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"qiyuan/hander"
)

func InitializeRoutes(router *gin.Engine) {

	//base requirements
	// Handle the index route
	router.GET("/", hander.ShowIndexPage)

	// use to create an order
	router.PUT("/orders", hander.HandleCreateOrder)

	// use to update an order
	router.POST("/orders", hander.HandleUpdateOrder)

	// use to get details about an order
	router.GET("/orders/view/:order_id", hander.HandleGetOrder)

	// use to get an order list by different condition
	router.GET("/orders", hander.HandleQueryOrders)

	// another requirements
	// use to get an order list by different condition
	router.GET("/orders/checkout", hander.HandleCheckOut)

	// use to get an order list by different condition
	router.POST("/file/upload/:order_id", hander.HandleUpload)

	// use to get an order list by different condition
	router.GET("/file/download", hander.HandleDownload)

	router.StaticFS("/files", http.Dir("../file"))
}

