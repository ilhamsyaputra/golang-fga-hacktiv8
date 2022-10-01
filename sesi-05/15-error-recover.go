// RECOVER, wajib pake keyword defer

package main

import (
	"errors"
	"fmt"
)

func main() {
	defer catchErr()

	var password string

	fmt.Scanln(&password)

	if valid, err := validatePassword(password); err != nil {
		panic(err.Error())
	} else {
		fmt.Println(valid)
	}
}

func catchErr() {
	if r := recover(); r != nil {
		fmt.Println("Error occured:", r)
	} else {
		fmt.Println("Application running successfully")
	}
}

func validatePassword(password string) (string, error) {
	pl := len(password)

	if pl < 5 {
		return "", errors.New("Password must be at least 4 characters long")
	}

	return "Valid password", nil
}
