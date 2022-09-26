package main

import (
	"fmt"
	"strings"
)

func main() {

	// Looping 1
	fmt.Println("Looping 1")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	fmt.Println(strings.Repeat("#", 20))

	// Looping 2
	i := 0
	fmt.Println("Looping 2")
	for i < 5 {
		fmt.Println(i)
		i++
	}
	fmt.Println(strings.Repeat("#", 20))

	// Looping 3
	j := 0

	fmt.Println("Looping 3")
	for {
		fmt.Println(j)
		j++

		if j == 5 {
			break
		}
	}
	fmt.Println(strings.Repeat("#", 20))

	// break, continue, label
	fmt.Println("Loop with break, continue and label")
loopWithLabel:
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		if i%2 == 0 {
			continue
		}
		if i == 3 {
			break loopWithLabel
		}
	}

}
