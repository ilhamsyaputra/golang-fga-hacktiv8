package main

func main() {
	var emptyInterface interface{}

	_ = emptyInterface

	emptyInterface = "M. Ilham Syaputra"
	emptyInterface = "Jambi"
	emptyInterface = 25
	emptyInterface = true
	emptyInterface = []string{"ilham", "jambi"}

	/*
		Empty interface dapat menampung data dengan tipe data apa saja.
	*/

	emptyInterface = 20
	emptyInterface += 1
	/*
		kode ini akan menghasilkan error saat compile, karena kita melakukan operasi pertambahan dari tipe integer + interface
		untuk menanggulanginya, ada di kode dibawah (ASSERTION)
	*/

	emptyInterface = 20
	if value, ok := emptyInterface.(int); ok == true {
		emptyInterface = value + 1
	}

}
