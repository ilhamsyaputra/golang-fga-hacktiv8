package main

import (
	"fmt"
	"net/http"
)

var PORT = ":8080"

func greet(w http.ResponseWriter, r *http.Request) {
	msg := "Hello Golang!"
	fmt.Fprintf(w, msg)
}

func main() {
	/*
		fungsi HandleFunc digunakan untuk routing
		HandleFunc(PATH_ROUTING, FUNCTION(http.ResponseWriter, http.Request))

		fungsi ListendAndServe() digunakan untuk menjalankan server aplikasi
		ListendAndServe(PORT, HANDLER)
	*/
	http.HandleFunc("/", greet)
	http.ListenAndServe(PORT, nil)
}
