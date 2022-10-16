package routers

import (
	"backend/controllers"
	"backend/middlewares"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.Use(middlewares.EnableCORS())
	router.GET("/api/weather", controllers.GetWeather)

	return router
}
