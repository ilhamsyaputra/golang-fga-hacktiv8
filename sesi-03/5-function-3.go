// VARIADIC FUNCTION

package main

import (
	"fmt"
)

func print(names ...string) []map[string]string {
	var result []map[string]string

	for index, value := range names {
		key := fmt.Sprintf("Student%d", index+1)
		temp := map[string]string{
			key: value,
		}
		result = append(result, temp)
	}

	return result
}

func main() {
	studentList := print("Airel", "Nanda", "Milo", "Schannel", "Marco")

	fmt.Println(studentList)
}
