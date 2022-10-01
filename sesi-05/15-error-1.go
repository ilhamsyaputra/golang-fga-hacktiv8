package main

import (
	"fmt"
	"strconv"
)

func main() {
	var number int
	var err error

	number, err = strconv.Atoi("1234g")

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(number)
	}
}
