// GIVING VALUE TO A STRUCT

package main

import "fmt"

type Employee struct {
	name     string
	age      int
	division string
}

func main() {
	var employee Employee

	employee.name = "Ilham"
	employee.age = 23
	employee.division = "Data Team"

	fmt.Println(employee.name)
	fmt.Println(employee.age)
	fmt.Println(employee.division)

	var employee2 = Employee{name: "Nadine", age: 1, division: "Data Team"}
	fmt.Println(employee)
	fmt.Println(employee2)
}
