package hander

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"qiyuan/model"
	"qiyuan/service"
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





func HandleCreateOrder(c *gin.Context){

}

func HandleUpdateOrder(c *gin.Context){

}

func HandleGetOrder(c *gin.Context)  {
	orderID := c.Param("order_id")

	// do something to Check if the article ID is valid
	if order, err := service.GetOrderByID(orderID); err == nil {
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

func HandleQueryOrders(c *gin.Context) {

}