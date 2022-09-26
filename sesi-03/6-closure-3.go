// CLOSURE AS A RETURN VALUE

package main

import (
	"fmt"
	"strings"
)

func findStudent(students []string) func(string) string {
	return func(s string) string {
		var student string
		var position int

		for index, value := range students {
			if strings.ToLower(value) == strings.ToLower(s) {
				student = value
				position = index
				break
			}
		}

		if student == "" {
			return fmt.Sprintf("%s doesn't exist!", s)
		}
		return fmt.Sprintf("We found %s at position %d", s, position)
	}
}

func main() {
	var studentList = []string{"Airell", "Nanda", "Mailo", "Schannel", "Marco"}

	var find = findStudent(studentList)

	fmt.Println(find("Nandaaa"))
}
