package main

import (
	"fmt"
	"strings"
)

func main() {

	// ARRAY 1 DIMENSI
	array1d := [3]string{"Apel", "Mangga", "Jeruk"}

	fmt.Println(strings.Repeat("#", 20))
	fmt.Println("Array 1 Dimensi")

	// loop through array elements
	for index, value := range array1d {
		fmt.Println(index, value)
	}

	fmt.Println(strings.Repeat("#", 20))

	// ARRAY 2 DIMENSI
	array2d := [2][3]string{{"Apel", "Mangga", "Jeruk"}, {"Rambutan", "Durian", "Nanas"}}

	for _, element := range array2d {
		for _, value := range element {
			fmt.Println(value)
		}
	}
	fmt.Println(strings.Repeat("#", 20))

	// test akses element dari array 2 dimensi
	fmt.Println(array2d[0][1])
}
