package main

import (
	"JSONWebToken/database"
	"JSONWebToken/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}
