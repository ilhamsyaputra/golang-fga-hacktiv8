package main

import "belajar-go/helpers"

func main() {
	helpers.Greet()
	// helpers.greet() 			akan error karena yang di export itu wajib huruf kapital diawal

	var person = helpers.Person{}
	person.InvokeGreet()
}
