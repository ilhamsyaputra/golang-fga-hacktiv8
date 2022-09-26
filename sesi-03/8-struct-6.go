// SLICE OF ANONYMOUS STRUCT

package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {

	var employee = []struct {
		person   Person
		division string
	}{
		{person: Person{name: "John", age: 21}, division: "Software Engineer"},
		{person: Person{name: "Mary", age: 20}, division: "Finance"},
	}

	fmt.Printf("%+v \n", employee)

	for _, value := range employee {
		fmt.Printf("%+v \n", value)
	}
}
