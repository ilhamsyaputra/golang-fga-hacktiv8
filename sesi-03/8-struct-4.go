// ANONYMOUS STRUCT

package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {

	// Anonymous function tanpa pengisian field
	var employee1 = struct {
		person   Person
		division string
	}{}
	employee1.division = "Finance"
	employee1.person = Person{name: "John", age: 21}

	// Anonymous function dengan pengisian field
	var employee2 = struct {
		person   Person
		division string
	}{
		division: "Engineering",
		person:   Person{name: "Doe", age: 21},
	}

	fmt.Printf("Employee1: %+v \n", employee1)
	fmt.Printf("Employee2: %+v \n", employee2)
}
