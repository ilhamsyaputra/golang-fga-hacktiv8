package helper

var database = []struct {
	noAbsen   int
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}{
	{
		noAbsen:   1,
		nama:      "Ilham",
		alamat:    "Jambi",
		pekerjaan: "Mahasiswa",
		alasan:    "Ingin menjadi backend engineer",
	},
	{
		noAbsen:   2,
		nama:      "Syaputra",
		alamat:    "Jakarta",
		pekerjaan: "Mahasiswa",
		alasan:    "Ingin belajar bahasa Go",
	},
	{
		noAbsen:   3,
		nama:      "Rizky",
		alamat:    "Wamena",
		pekerjaan: "Mahasiswa",
		alasan:    "Ingin menambah pengetahuan mengenai Golang",
	},
}
