package main

import (
	"fmt"
)

func deferFunc() {
	fmt.Println("Defer function starts to execute")
}

func callDeferFunc() {
	defer deferFunc()
}

func main() {
	callDeferFunc()
	fmt.Println("Hi everyone!")
	// os.exit(1)
}

/*
	"Defer function starts to execute" akan tercetak duluan, karena fungsi deferFunc() dipanggil dengan keyword defer di fungsi callDeferFunc() dimana pada fungsi
	callDeferFunc() tidak terdapat operasi lain.
*/
