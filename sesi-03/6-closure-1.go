// DECLARE CLOSURE IN VARIABLE

package main

import "fmt"

func main() {
	evenNumbers := func(numbers ...int) []int {
		var result []int

		for _, value := range numbers {
			if value%2 == 0 {
				result = append(result, value)
			}
		}

		return result
	}

	var numbers = []int{4, 92, 12, 42, 90, 13}

	fmt.Println(evenNumbers(numbers...))
}
