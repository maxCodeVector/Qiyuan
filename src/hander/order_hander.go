package hander

import (
	"github.com/gin-gonic/gin"
	"model"
	"net/http"
)


// Render one of HTML, JSON or CSV based on the 'Accept' header of the request
// If the header doesn't specify this, HTML is rendered, provided that
// the template name is present
func render(c *gin.Context, data gin.H, templateName string) {

	switch c.Request.Header.Get("Accept") {
	case "application/json":
		// Respond with JSON
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		// Respond with XML
		c.XML(http.StatusOK, data["payload"])
	default:
		// Respond with HTML
		c.HTML(http.StatusOK, templateName, data)
	}

}


func ShowIndexPage(c *gin.Context) {
	orders := model.GetOrders()

	render(
		c,
		// Pass the data that the page uses (in this case, 'title')
		gin.H{
			"title":   "Home Page",
			"payload": *orders,
		},
		"index.html",
	)
}

func GetOrder(c *gin.Context)  {
	orderID := c.Param("order_id")

	// do something to Check if the article ID is valid
	if order, err := model.GetOrderByID(orderID); err == nil {
		// Call the relative method of the Context to render a template
		render(
			c,
			// Pass the data that the page uses (in this case, 'title')
			gin.H{
				"title":   order.UserName,
				"payload": order,
			},
			"order.html",
		)

	} else {
		// If the article is not found, abort with an error
		c.AbortWithError(http.StatusNotFound, err)
	}
}

func HandleVerification (c *gin.Context)  {
	if c.Request.Method == "OPTIONS" {
		// setup headers
		c.Header("Allow", "POST, GET, OPTIONS")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "origin, content-type, accept")
		c.Header("Content-Type", "application/json")
		c.Status(http.StatusOK)
	} else if c.Request.Method == "POST" {
		var u model.User
		c.BindJSON(&u)
		c.JSON(http.StatusOK, gin.H{
			"user": u.Username,
			"pass": u.Password,
		})
		print(u.Password, u.Username)
	}

}



func HandleCreateOrder(c *gin.Context){

}

func HandleUpdateOrder(c *gin.Context){

}

func HandleGetOrder(c *gin.Context)  {

}

func HandleQueryOrders(c *gin.Context) {

}