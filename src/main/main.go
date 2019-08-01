package main

import (
	"github.com/gin-gonic/gin"
	myr "router"
)

func main()  {
	//db.Main()
	router := gin.Default()
	router.LoadHTMLGlob("src/templates/*")

	myr.InitializeRoutes(router)

	router.Run()


}