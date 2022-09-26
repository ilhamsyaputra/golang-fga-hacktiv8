package main

import (
	"fmt"
)

func main() {
	person := map[string]string{
		"Name":        "M. Ilham Syaputra",
		"Age":         "23",
		"Nationality": "Indonesian",
	}
	fmt.Printf("Hi! you must be %s and you are %s years old\n", person["Name"], person["Age"])

	// LOOPING WITH MAP
	for key, value := range person {
		fmt.Println(key, ":", value)
	}

	// DELETING ELEMENT
	fmt.Println("Before deleting:", person)

	delete(person, "Age")
	fmt.Println("After deleting:", person)

	// DETECTING VALUE
	value2, isExist := person["Name"]
	fmt.Println(value2, isExist)

	value3, isExist := person["Age"]
	fmt.Println(value3, isExist)
}
