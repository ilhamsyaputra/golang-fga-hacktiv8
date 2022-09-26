// EMBEDDED STRUCT

package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	division string
	person   Person
}

func main() {
	var employee1 = Employee{}

	employee1.division = "Data"
	employee1.person.name = "John"
	employee1.person.age = 21

	fmt.Printf("%+v", employee1)
}
