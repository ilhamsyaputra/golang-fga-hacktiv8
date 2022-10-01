/*
	channels adalah mekanisme untuk goroutine dapat saling berkomunikasi. maksud berkomunikasi adalah proses pertukaran data antar goroutine
*/

package main

import "fmt"

func main() {
	c := make(chan string)

	go introduce("Ilham", c)
	go introduce("Kiki", c)
	go introduce("Syifa", c)
	go introduce("Nadine", c)

	msg1 := <-c
	fmt.Println(msg1)

	msg2 := <-c
	fmt.Println(msg2)

	msg3 := <-c
	fmt.Println(msg3)

	msg4 := <-c
	fmt.Println(msg4)

	close(c)
}

func introduce(student string, c chan string) {
	result := fmt.Sprintln("Hi! My name is", student)

	c <- result
}
