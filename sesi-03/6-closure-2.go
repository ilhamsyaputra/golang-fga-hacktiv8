// CLOSURE IIFE (IMMEDIATELY-INVOKED FUNCTION EXPRESSION)

package main

import "fmt"

func main() {
	var isPalindrome = func(word string) bool {
		var temp string = ""

		for i := len(word) - 1; i >= 0; i-- {
			temp += string(byte(word[i]))
		}

		return temp == word
	}("katak")

	fmt.Println(isPalindrome)
}
