package main

import (
	"github.com/gin-gonic/gin"
	myrounter "qiyuan/router"
)

func main()  {

	router := gin.Default()
	router.LoadHTMLGlob("../templates/*")

	myrounter.InitializeRoutes(router)

	err := router.Run()
	if err != nil {
		print(err.Error())
	}
}