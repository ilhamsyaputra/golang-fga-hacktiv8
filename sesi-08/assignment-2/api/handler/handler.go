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

	if allOrders.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"statusCode": http.StatusInternalServerError,
			"message":    "Error occured!",
		})
	} else if allOrders.RowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "No Data",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "Success retrieve orders data",
			"data":    orders,
		})
	}

}

func UpdateOrderHandler(ctx *gin.Context) {
	db := database.GetDB()
	orderId, _ := strconv.Atoi(ctx.Param("orderId"))
	updatedOrder := models.Order{}
	updatedItem := models.Item{}

	if err := ctx.ShouldBindJSON(&updatedOrder); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	updatedOrder.Items[0].OrderID = uint(orderId)
	updatedOrder.ID = uint(orderId)

	err := db.Model(&updatedOrder).Where("order_id = ?", orderId).Updates(
		updatedOrder,
	).Error

	if err != nil {
		fmt.Println("Error updating product:", err.Error())
		return
	} else {
		db.Model(&updatedItem).Where("order_id = ?", orderId).Updates(
			models.Item{
				ItemCode:    updatedOrder.Items[0].ItemCode,
				Description: updatedOrder.Items[0].Description,
				Quantity:    updatedOrder.Items[0].Quantity,
			},
		)
		ctx.JSON(http.StatusOK, gin.H{
			"statusCode": http.StatusOK,
			"message":    "success update order data",
		})
	}
}

func DeleteOrderHandler(ctx *gin.Context) {
	db := database.GetDB()
	orderId, _ := strconv.Atoi(ctx.Param("orderId"))
	item := models.Item{}
	order := models.Order{}

	deleteItemData := db.Where("order_id = ?", orderId).Delete(&item)

	if deleteItemData.Error != nil {
		fmt.Println("Error deleting product:", deleteItemData.Error.Error())
		return
	} else if deleteItemData.RowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"statusCode": http.StatusNotFound,
			"message":    "data not found",
		})
	} else {
		db.Where("order_id = ?", orderId).Delete(&order)
		ctx.JSON(http.StatusOK, gin.H{
			"statusCode": http.StatusOK,
			"message":    "success delete order data",
		})
	}
}
