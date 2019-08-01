package main

import (
	"github.com/gin-gonic/gin"
	myrounter "router"
)

func main()  {

	router := gin.Default()
	router.LoadHTMLGlob("src/templates/*")

	myrounter.InitializeRoutes(router)

	router.Run()


}