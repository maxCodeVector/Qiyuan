package router

import (
	"github.com/gin-gonic/gin"
	"hander"
)

func InitializeRoutes(router *gin.Engine) {

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

}

