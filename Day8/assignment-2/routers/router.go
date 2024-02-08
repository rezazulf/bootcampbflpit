package routers

import (
	"assignment-2_reza/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer(c *controllers.Controller) *gin.Engine {
	router := gin.Default()

	router.POST("/orders", c.CreateOrder)
	router.GET("/orders", c.GetOrders)
	router.GET("/orders/:orderId", c.GetOrderByID)
	router.PUT("/orders/:orderId", c.UpdateOrder)
	router.DELETE("/orders/:orderId", c.DeleteOrder)

	return router
}
