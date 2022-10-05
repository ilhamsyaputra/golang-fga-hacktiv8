package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var jsonString = `
		{
			"full_name": "M. Ilham Syaputra",
			"email": "ilhamsyaputra25@gmail.com",
			"age": 23
		}
	`

	var result map[string]interface{}

	var err = json.Unmarshal([]byte(jsonString), &result)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("%+v\n", result)
}
