/*
M. Ilham Syaputra
149368582101-12
*/

package main

import (
	"assignment-1/helper"
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) == 1 {
		fmt.Println("Pastikan anda memasukkan nomor absen")
		os.Exit(1)
	}

	noAbsen, err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Println("Masukkan nomor absen yang benar")
		return
	}

	helper.GetData(noAbsen)
}
