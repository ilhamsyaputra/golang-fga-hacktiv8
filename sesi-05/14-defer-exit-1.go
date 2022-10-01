/*
	defer adalah keyword untuk memanggil function yang dimana proses eksekusinya ditahan hingga block sebuah function selesai.
*/

package main

import "fmt"

func main() {
	defer fmt.Println("defer function starts to execute")
	fmt.Println("Hi everyone!")
	fmt.Println("My name is Ilham!")
}
