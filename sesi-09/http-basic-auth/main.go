/*
	untuk menjalankan: go run *.go

	curl -X GET --user batman:secret123 http://localhost:9000/student
	curl -X GET --user batman:secret123 http://localhost:9000/student?id=s001
*/

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ActionStudent(w http.ResponseWriter, r *http.Request) {
	if !Auth(w, r) {
		return
	}
	if !AllowOnlyGet(w, r) {
		return
	}

	if id := r.URL.Query().Get("id"); id != "" {
		OutputJSON(w, SelectedStudent(id))
		return
	}

	OutputJSON(w, GetStudents())
}

func OutputJSON(w http.ResponseWriter, o interface{}) {
	res, err := json.Marshal(o)

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func main() {
	http.HandleFunc("/student", ActionStudent)

	server := new(http.Server)
	server.Addr = ":9000"

	fmt.Println("Server running at http://localhost:9000/")
	server.ListenAndServe()
}
