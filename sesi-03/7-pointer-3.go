// POINTER AS PARAMETER

package main

import "fmt"

func changeValue(number *int) {
	*number = 20
}

func main() {
	var a int = 10

	fmt.Println("Before:", a)

	changeValue(&a)

	fmt.Println("After:", a)
}
