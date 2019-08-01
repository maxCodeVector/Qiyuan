package main

import (
	"github.com/gin-gonic/gin"
	myrounter "router"
)

func main()  {
	//orders := model.GetOrders()
	//print((*orders)[0].FileUrl)

	router := gin.Default()
	router.LoadHTMLGlob("src/templates/*")

	myrounter.InitializeRoutes(router)

	router.Run()


}