// POINTER TO A STRUCT

package main

import (
	"fmt"
	"strings"
)

type Employee struct {
	name     string
	age      int
	division string
}

func main() {
	employee1 := Employee{name: "John", age: 21, division: "Data"}

	var employee2 *Employee = &employee1

	fmt.Println("Employee1 name:", employee1.name)
	fmt.Println("Employee2 name:", employee2.name)

	fmt.Println(strings.Repeat("#", 20))

	employee2.name = "Doe"
	fmt.Println("Employee1 name:", employee1.name)
	fmt.Println("Employee2 name:", employee2.name)
}
