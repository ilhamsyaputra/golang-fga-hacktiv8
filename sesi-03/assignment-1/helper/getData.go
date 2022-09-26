package helper

import (
	"fmt"
	"os"
)

func PrintPersonInfo(absen int) {
	person := database[absen-1]
	fmt.Printf("Nama: %s \n", person.nama)
	fmt.Printf("Alamat: %s \n", person.alamat)
	fmt.Printf("Pekerjaan: %s \n", person.pekerjaan)
	fmt.Printf("Alasan: %s \n", person.alasan)
}

func GetData(absen int) {
	if absen <= 0 || absen > len(database) {
		fmt.Println("Masukkan nomor absen yang benar")
		os.Exit(1)
	}
	PrintPersonInfo(absen)
}
