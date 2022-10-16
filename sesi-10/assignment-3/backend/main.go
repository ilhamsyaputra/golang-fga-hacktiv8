package main

import "backend/routers"

func main() {
	routers.StartServer().Run(":8080")
}
