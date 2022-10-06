package main

import (
	"assignment-2/api/routes"
	"assignment-2/database"
)

func main() {
	database.StartDB()

	var PORT = ":8080"
	routes.StartServer().Run(PORT)
}
