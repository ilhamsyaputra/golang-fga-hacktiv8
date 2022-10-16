package controllers

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Weather struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

func generateRandomNumber() int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(100-1) + 1
}

func GetWeather(ctx *gin.Context) {
	var weather Weather

	if ctx.Request.Method != "GET" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Only GET request allowed!",
		})
	}

	weather.Water = generateRandomNumber()
	weather.Wind = generateRandomNumber()

	ctx.JSON(http.StatusOK, gin.H{
		"status": weather,
	})
}
