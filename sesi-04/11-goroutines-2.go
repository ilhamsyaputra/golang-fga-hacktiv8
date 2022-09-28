// GOROUTINES adalah cara untuk mengeksekusi kode secara asynchronous

package main

import (
	"fmt"
	"runtime"
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
	fmt.Println("Main execution finished")
}
