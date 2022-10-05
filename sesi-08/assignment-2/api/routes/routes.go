package routes

import (
	"assignment-2/api/handler"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/orders", handler.CreateOrderHandler)

	return router
}
