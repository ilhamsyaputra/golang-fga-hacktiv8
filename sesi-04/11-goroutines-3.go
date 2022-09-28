/*
	Goroutines membutuhkan waktu yang sedikit lebih lama untuk memulai dibandingkan dengan function biasa.
	Untuk itu, ditambahkan function time.Sleep() yang akan menahan function main() untuk langsung menyelesaikan eksekusinya LINE 37
	Namun cara time.Sleep() kurang baik, lebih baik menggunakan waitgroup
*/

package main

import (
	"fmt"
	"runtime"
	"time"
)

func firstProcess(index int) {
	fmt.Println("First process function started")
	for i := 1; i <= index; i++ {
		fmt.Println("i =", index)
	}
	fmt.Println("First process function ended")
}

func secondProcess(index int) {
	fmt.Println("Second process function started")
	for i := 1; i <= index; i++ {
		fmt.Println("i =", index)
	}
	fmt.Println("Second process function ended")
}

func main() {
	fmt.Println("Main execution started")

	go firstProcess(7)
	secondProcess(7)

	fmt.Println("# of Goroutines:", runtime.NumGoroutine())
	time.Sleep(time.Second * 2)
	fmt.Println("Main execution finished")
}
