// SLICE OF STRUCT

package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {

	var people = []Person{
		{name: "John", age: 21},
		{name: "Mary", age: 21},
		{name: "Peter", age: 21},
	}

	fmt.Printf("%+v\n", people)
	fmt.Println()

	for _, person := range people {
		fmt.Printf("%+v\n", person)
	}
}
