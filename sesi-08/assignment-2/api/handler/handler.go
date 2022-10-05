package handler

import (
	"assignment-2/database"
	"assignment-2/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOrderHandler(ctx *gin.Context) {
	db := database.GetDB()
	newOrder := models.Order{}

	if err := ctx.ShouldBindJSON(&newOrder); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err := db.Create(&newOrder).Error

	if err != nil {
		fmt.Println("Error creating order.", err)
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": fmt.Sprintf("Order created!"),
	})
}

func GetOrderHandler(ctx *gin.Context) {
	db := database.GetDB()
	orders := models.Order{}

	err := db.Preload("Items").Find(&orders).Error

	if err != nil {
		fmt.Println("Error getting orders", err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success retrieve orders data",
		"order":   orders,
	})
}
