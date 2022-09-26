package main

import "fmt"

func main() {
	lang := "jp"

	if lang == "jp" {
		fmt.Println("Ohayou gozaimasu")
	} else if lang == "id" {
		fmt.Println("Halo! Selamat pagi")
	} else {
		fmt.Println("We don't recognize your language!")
	}
}
