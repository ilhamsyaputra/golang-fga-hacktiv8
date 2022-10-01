package main

import "fmt"

func main() {
	c := make(chan string)

	students := []string{"Ilham", "Kiki", "Syifa", "Nadine"}

	for _, value := range students {
		go func(student string) {
			fmt.Println("Student: ", student)
			result := fmt.Sprintf("Hi, My name is %s", student)
			c <- result
		}(value)
	}

	for i := 1; i <= 4; i++ {
		print(c)
	}
	close(c)
}

func print(c chan string) {
	fmt.Println(<-c)
}
