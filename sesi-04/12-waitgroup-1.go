/*
	Waitgroup adalah struct dari package sync, yang digunakan untuk melakukan sinkronisasi terhadap goroutines.

	1. Line 28: Menggunakan method Add() untuk menambah counter dari waitgroup, untuk memberitahu waitgroup tentang jumlah goroutines. terdapat 4 goroutines
		yang harus ditunggu sebelum function main menghentikan eksekusinya
	2. Line 32: Wait(), menunggu seluruh goroutine menyelesaikan prosesnya, method Wait() menahan function main hingga seluruh goroutine selesai
	3. Line 19: Done(), Untuk memberitahukan waitgroup tentang goroutine yang telah menyelesaikan prosesnya.
*/

package main

import (
	"fmt"
	"sync"
)

func printFruit(index int, fruit string, wg *sync.WaitGroup) {
	fmt.Printf("Index => %d, fruit => %s\n", index, fruit)
	wg.Done()
}

func main() {
	fruits := []string{"Apple", "Mango", "Durian", "Banana", "Rambutan", "PineApple"}

	var wg sync.WaitGroup

	for index, fruit := range fruits {
		wg.Add(1)
		go printFruit(index, fruit, &wg)
	}

	wg.Wait()
}
