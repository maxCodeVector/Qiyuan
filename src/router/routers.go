package router

import (
	"github.com/gin-gonic/gin"
	"hander"
)

func InitializeRoutes(router *gin.Engine) {

	// Handle the index route
	router.GET("/", hander.ShowIndexPage)

	// Handle GET requests at /article/view/some_article_id
	router.GET("/order/view/:order_id", hander.GetOrder)

	router.POST("/api", hander.HandleVerification)
	router.OPTIONS("/api", hander.HandleVerification)
}

