package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	em "test/models"
)

var PORT = ":8080"

func main() {
	http.HandleFunc("/employees", getEmployees)
	http.HandleFunc("/createemployee", createEmployee)

	fmt.Println("Application running on port ", PORT)
	http.ListenAndServe(PORT, nil)
}

func getEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		json.NewEncoder(w).Encode(em.Employees)
		return
	}
	http.Error(w, "Invalid Method", http.StatusBadRequest)
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
		newEmployee := em.Employee{
			ID:       len(em.Employees) + 1,
			Name:     name,
			Age:      convertAge,
			Division: division,
		}

		em.Employees = append(em.Employees, newEmployee)

		json.NewEncoder(w).Encode(newEmployee)
		return
	}
	http.Error(w, "Invalid Method", http.StatusBadRequest)
}
