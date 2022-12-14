// POINTER (MEMORY ADDRESS)
package main

import "fmt"

func main() {
	var firstNumber int = 4
	var secondNumber *int = &firstNumber

	fmt.Println("firstNumber value:", firstNumber)
	fmt.Println("firstNumber memory address:", secondNumber)

	fmt.Println("secondNumber value:", *secondNumber)
	fmt.Println("secondNumber memory address", secondNumber)
}
