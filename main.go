package main

import (
	"fmt"
	em "go_crud/employee"
	handler "go_crud/http"
	"log"
	"net/http"
)

func main() {
	// Create a new instance of EmployeeStore
	empStore := em.NewEmployeeStore()

	// Test data
	emp1 := em.Employee{ID: 1, Name: "Aadit Patil", Position: "Software Engineer 2", Salary: 150000}
	emp2 := em.Employee{ID: 2, Name: "Rurui Jushi", Position: "Data Scientist", Salary: 60000}
	emp3 := em.Employee{ID: 3, Name: "Manmohan Ramindar", Position: "Product Manager", Salary: 70000}

	// Add test data to store
	empStore.CreateEmployee(emp1)
	empStore.CreateEmployee(emp2)
	empStore.CreateEmployee(emp3)
	hand := handler.Handler{
		Es: empStore,
	}
	// HTTP endpoints
	http.HandleFunc("/employees", hand.ListEmployees)

	// Start server
	fmt.Println("Server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
