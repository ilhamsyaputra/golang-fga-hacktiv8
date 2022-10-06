package handler

import (
	"assignment-2/database"
	"assignment-2/models"
	"fmt"
	"net/http"
	"strconv"

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
	orders := []models.Order{}

	allOrders := db.Preload("Items").Find(&orders)

	if allOrders.RowsAffected != 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "Success retrieve orders data",
			"order":   orders,
		})
	} else {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "No Data",
		})
	}

}

func UpdateOrderHandler(ctx *gin.Context) {
	db := database.GetDB()
	orderId, _ := strconv.Atoi(ctx.Param("orderId"))
	updatedOrder := models.Order{}

	if err := ctx.ShouldBindJSON(&updatedOrder); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	updatedOrder.Items[0].OrderID = uint(orderId)

	err := db.Model(&updatedOrder).Where("order_id = ?", orderId).Updates(
		updatedOrder,
	).Error

	if err != nil {
		fmt.Println("Error updating product:", err.Error())
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "success update order data",
		})
	}
}

func DeleteOrderHandler(ctx *gin.Context) {
	db := database.GetDB()
	orderId, _ := strconv.Atoi(ctx.Param("orderId"))
	order := models.Item{}

	err := db.Where("order_id = ?", orderId).Delete(&order).Error

	if err != nil {
		fmt.Println("Error deleting product:", err.Error())
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "success delete order data",
		})
	}

}
