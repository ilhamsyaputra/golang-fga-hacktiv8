package main

import (
	"fmt"
	"reflect"
)

func main() {
	number := 25
	reflectValue := reflect.ValueOf(number)

	fmt.Println("Variable number .ValueOf():", reflect.ValueOf(number))
	fmt.Println("Variable number .ValueOf().Type():", reflect.ValueOf(number).Type())
	fmt.Println("Variable number TypeOf().String:", reflect.TypeOf(number).String())

	// Cek tipe data
	fmt.Println(reflectValue.Type())
	if reflectValue.Kind() == reflect.Int {
		fmt.Println("Nilai variable number:", reflectValue.Int())
	}

	// daftar tipe data: https://pkg.go.dev/reflect#Kind

	// ACCESING VALUE USING INTERFACE METHOD
	fmt.Println()
	fmt.Println("Accessing value using interface method:")
	fmt.Println(reflectValue.Type())
	fmt.Println(reflectValue.Interface())
	fmt.Println(reflectValue.Interface().(int))
}
