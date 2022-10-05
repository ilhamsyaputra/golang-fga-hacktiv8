package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}

func main() {
	var jsonString = `[
		{
			"full_name": "M. Ilham Syaputra",
			"email": "ilhamsyaputra25@gmail.com",
			"age": 23
		},
		{
			"full_name": "Nadine",
			"email": "nadinealyssaazzahra@gmail.com",
			"age": 0
		}
	]`

	var result []Employee

	var err = json.Unmarshal([]byte(jsonString), &result)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for index, value := range result {
		fmt.Printf("Index %d: %+v \n", index+1, value)
	}
}
