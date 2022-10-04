package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

type Employee struct {
	ID       int
	Name     string
	Age      int
	Division string
}

var employees = []Employee{
	{ID: 1, Name: "Ilham", Age: 23, Division: "IT"},
	{ID: 2, Name: "Rizky", Age: 20, Division: "Legal"},
	{ID: 3, Name: "Syifa", Age: 11, Division: "Marketing"},
	{ID: 4, Name: "Nadine", Age: 0, Division: "Finance"},
}

var PORT = ":8080"

func getEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		template, err := template.ParseFiles("api-template.html")

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		template.Execute(w, employees)
		return
	}

	http.Error(w, "Invalid Method", http.StatusBadRequest) // Jika bukan method GET akan error
}

func createEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		name := r.FormValue("name")
		age := r.FormValue("age")
		division := r.FormValue("division")

		convertAge, err := strconv.Atoi(age)

		if err != nil {
			http.Error(w, "Invalid age", http.StatusBadRequest)
			return
		}

		newEmployee := Employee{
			ID:       len(employees) + 1,
			Name:     name,
			Age:      convertAge,
			Division: division,
		}

		employees = append(employees, newEmployee)

		json.NewEncoder(w).Encode(newEmployee)
		return
	}

	http.Error(w, "Invalid method", http.StatusBadRequest)
}

func main() {
	http.HandleFunc("/employees", getEmployees)
	http.HandleFunc("/employee", createEmployee)

	fmt.Println("Application is running on port", PORT)
	http.ListenAndServe(PORT, nil)
}
