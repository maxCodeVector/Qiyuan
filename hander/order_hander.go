package hander

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"qiyuan/model"
	"qiyuan/service"
	"strconv"
	"time"
)

// Render one of HTML, JSON or CSV based on the 'Accept' header of the request
// If the header doesn't specify this, HTML is rendered, provided that
// the template name is present
func render(c *gin.Context, data gin.H, status int, templateName string) {

	switch c.Request.Header.Get("Accept") {
	case "application/json":
		// Respond with JSON
		c.JSON(status, data["payload"])
	case "application/xml":
		// Respond with XML
		c.XML(status, data["payload"])
	default:
		// Respond with HTML
		c.HTML(status, templateName, data)
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
		http.StatusOK,
		"index.html",
	)
}

func HandleCreateOrder(c *gin.Context) {
	userName := c.Query("userName")
	status := c.Query("status")
	fileUrl := c.Query("fileUrl")

	amount, err := strconv.ParseFloat(c.Query("amount"), 64)
	if err != nil {
		render(
			c,
			// Pass the data that the page uses (in this case, 'title')
			gin.H{
				"title":   "not float amount",
				"payload": "What's your problem?",
			},
			http.StatusBadRequest,
			"error.html",
		)
		return
	}

	timeUnix := time.Now().Unix()

	order := model.Order{UserName: userName, Amount: amount, FileUrl: fileUrl, Status: status, OrderId: strconv.FormatInt(timeUnix, 10)}
	service.CreateOrder(&order)
	render(
		c,
		// Pass the data that the page uses (in this case, 'title')
		gin.H{
			"title":   "Success",
			"payload": "Create Success",
		},
		http.StatusOK,
		"error.html",
	)

}

func HandleUpdateOrder(c *gin.Context) {

	orderId := c.Query("orderId")
	status := c.Query("status")
	fileUrl := c.Query("fileUrl")

	amount, err := strconv.ParseFloat(c.Query("amount"), 64)
	if err != nil {
		render(
			c,
			// Pass the data that the page uses (in this case, 'title')
			gin.H{
				"title":   "not float amount",
				"payload": "What's your problem",
			},
			http.StatusBadRequest,
			"error.html",
		)
		return
	}

	order := model.Order{Amount: amount, Status: status, FileUrl: fileUrl, OrderId: orderId}
	if service.UpdateOrder(&order) {
		render(
			c,
			// Pass the data that the page uses (in this case, 'title')
			gin.H{
				"title":   "Success",
				"payload": "Update Success",
			},
			http.StatusOK,
			"error.html",
		)
	} else {
		render(
			c,
			// Pass the data that the page uses (in this case, 'title')
			gin.H{
				"title":   "Fail",
				"payload": "Update fail",
			},
			http.StatusNotFound,
			"error.html",
		)
	}

}

func HandleGetOrder(c *gin.Context) {
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
			http.StatusOK,
			"order.html",
		)

	} else {
		// If the article is not found, abort with an error
		c.AbortWithError(http.StatusNotFound, err)
	}
}

func HandleQueryOrders(c *gin.Context) {
	amount := c.Query("amount")
	time := c.Query("time")
	userName := c.Query("userName")
	orderBytime, err := strconv.ParseBool(time)
	if err != nil {
		orderBytime = false
	}

	orderByAmount, err2 := strconv.ParseBool(amount)
	if err2 != nil {
		orderByAmount = false
	}

	orders := service.FuzzySearchOrder(userName, orderBytime, orderByAmount)
	render(
		c,
		// Pass the data that the page uses (in this case, 'title')
		gin.H{
			"title":   "Home Page",
			"payload": *orders,
		},
		http.StatusOK,
		"index.html",
	)
}

func HandleCheckOut(c *gin.Context) {
	userName := c.PostForm("userName")
	totalAmount := service.CheckOut(userName)

	render(
		c,
		// Pass the data that the page uses (in this case, 'title')
		gin.H{
			"title":   "You totally checkout amount",
			"payload": totalAmount,
		},
		http.StatusOK,
		"error.html",
	)
}
