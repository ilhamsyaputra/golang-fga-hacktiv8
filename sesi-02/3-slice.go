package main

import (
	"fmt"
	"strings"
)

func main() {

	// MEMBUAT SLICE CARA 1
	buah := []string{"Mangga", "Jeruk", "Apel"}

	fmt.Println(strings.Repeat("#", 20))
	fmt.Println("SLICE CARA 1")
	fmt.Println(buah)
	fmt.Println(strings.Repeat("#", 20))

	// MEMBUAT SLICE CARA 2
	buah2 := make([]string, 3)
	fmt.Println("Slice dari buah2", buah2)

	buah2[0] = "Durian"
	fmt.Println("Slice dari buah2 setelah diisi elemennya", buah2)

	// MENAMBAHKAN ELEMEN PADA SLICE DENGAN APPEND
	buah2 = append(buah2, "Mengkudu", "Rambutan", "Duku", "Salak")
	fmt.Println("Slice dari buah2 setelah diisi dengan append", buah2) // [Durian   Mengkudu Rambutan Duku Salak], elemen indeks ke 2 dan 3 adalah string kosong ""

	// COPY SLICE
	nn := copy(buah2, buah)
	fmt.Println(buah2)
	fmt.Println(nn)
}
