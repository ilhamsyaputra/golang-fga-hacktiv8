package routes

import (
	"assignment-2/api/handler"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/orders", handler.CreateOrderHandler)
	router.GET("/orders", handler.GetOrderHandler)
	router.PUT("/orders/:orderId", handler.UpdateOrderHandler)
	router.DELETE("/orders/:orderId", handler.DeleteOrderHandler)

	return router
}
