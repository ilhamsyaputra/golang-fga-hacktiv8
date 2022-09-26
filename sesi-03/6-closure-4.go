// CLOSURE CALLBACK

package main

import "fmt"

func findOddNumbers(numbers []int, callback func(int) bool) int {
	var totalOddNumbers int

	for _, value := range numbers {
		if callback(value) {
			totalOddNumbers++
		}
	}
	return totalOddNumbers
}

func main() {
	var numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var find = findOddNumbers(numbers, func(number int) bool {
		return number%2 == 0
	})

	fmt.Println("Total odd number:", find)
}
