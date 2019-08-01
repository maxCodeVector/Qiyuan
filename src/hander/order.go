package hander

import (
	"github.com/gin-gonic/gin"
	"model"
	"net/http"
)

func ShowIndexPage(c *gin.Context) {
	orders := model.GetOrders()

	// Call the HTML method of the Context to render a template
	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		"index.html",
		// Pass the data that the page uses (in this case, 'title')
		gin.H{
			"title":   "Home Page",
			"payload": *orders,
		},
	)
}

func GetOrder(c *gin.Context)  {
	orderID := c.Param("order_id")

	// do something to Check if the article ID is valid

	if order, err := model.GetOrderByID(orderID); err == nil {
		// Call the HTML method of the Context to render a template
		c.HTML(
			// Set the HTTP status to 200 (OK)
			http.StatusOK,
			// Use the index.html template
			"order.html",
			// Pass the data that the page uses
			gin.H{
				"title":   order.UserName,
				"payload": order,
			},
		)

	} else {
		// If the article is not found, abort with an error
		c.AbortWithError(http.StatusNotFound, err)
	}

}