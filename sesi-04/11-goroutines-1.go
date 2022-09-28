package main

import "fmt"

func goroutine() {
	fmt.Println("Go Routines test")
}

func main() {
	go goroutine()
}
