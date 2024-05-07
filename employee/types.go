package employee

import "sync"

// Employee struct defines the structure of an employee
type Employee struct {
	ID       int
	Name     string
	Position string
	Salary   float64
}

// EmployeeStore is an in-memory store for employees
type EmployeeStore struct {
	sync.RWMutex
	Employees map[int]Employee
}
