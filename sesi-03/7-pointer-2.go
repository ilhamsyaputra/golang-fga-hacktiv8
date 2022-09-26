// changing value through pointer

package main

import (
	"fmt"
	"strings"
)

func main() {
	var firstPerson string = "John"
	var secondPerson *string = &firstPerson

	fmt.Println("firstPerson value:", firstPerson)
	fmt.Println("firstPerson memory address:", &firstPerson)
	fmt.Println("secondPerson value:", *secondPerson)
	fmt.Println("secondPerson memory address:", secondPerson)

	fmt.Println(strings.Repeat("#", 20))
	fmt.Println()

	*secondPerson = "Doe"

	fmt.Println("firstPerson value:", firstPerson)
	fmt.Println("firstPerson memory address:", &firstPerson)
	fmt.Println("secondPerson value:", *secondPerson)
	fmt.Println("secondPerson memory address:", secondPerson)
}
