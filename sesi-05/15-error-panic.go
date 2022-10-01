// PANIC

package main

import (
	"errors"
	"fmt"
)

func main() {
	var password string

	fmt.Scanln(&password)

	if valid, err := validatePassword(password); err != nil {
		panic(err.Error())
	} else {
		fmt.Println(valid)
	}
}

func validatePassword(password string) (string, error) {
	pl := len(password)

	if pl < 5 {
		return "", errors.New("Password must be at least 4 characters long")
	}

	return "Valid password", nil
}
