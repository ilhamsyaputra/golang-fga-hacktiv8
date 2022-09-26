// VARIADIC FUNCTION WITH VARIADIC PARAMETER

package main

import "fmt"

func sum(numbers ...int) int {
	total := 0

	for _, value := range numbers {
		total += value
	}
	return total
}

func main() {
	numberLists := []int{1, 2, 3, 4, 5, 6, 7}

	result := sum(numberLists...)
	fmt.Println("Result:", result)
}
