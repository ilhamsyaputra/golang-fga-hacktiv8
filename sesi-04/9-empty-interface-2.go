package main

import "fmt"

func main() {
	// membuat sebuah map dengan tipe data value yang berbeda dengan memanfaatkan empty interface

	mapVariousDataTypeValue := map[string]interface{}{
		"name":       "ilham",
		"age":        23,
		"unemployed": true,
	}

	fmt.Println(mapVariousDataTypeValue)

	// membuat slice dengan tipe data value yang berbeda dengan memanfaatkan empty interface

	sliceVariousDataTypeValue := []interface{}{"Ilham", 23, true}
	fmt.Println(sliceVariousDataTypeValue)
}
