package employee

import "fmt"

// NewEmployeeStore creates a new instance of EmployeeStore
func NewEmployeeStore() *EmployeeStore {
	return &EmployeeStore{
		Employees: make(map[int]Employee),
	}
}

// CreateEmployee adds a new employee to the store
func (es *EmployeeStore) CreateEmployee(emp Employee) {
	es.Lock()
	defer es.Unlock()
	es.Employees[emp.ID] = emp
}

// GetEmployeeByID retrieves an employee from the store by ID
func (es *EmployeeStore) GetEmployeeByID(id int) (Employee, error) {
	es.RLock()
	defer es.RUnlock()
	emp, ok := es.Employees[id]
	if !ok {
		return Employee{}, fmt.Errorf("employee with ID %d not found", id)
	}
	return emp, nil
}

// UpdateEmployee updates the details of an existing employee
func (es *EmployeeStore) UpdateEmployee(id int, emp Employee) error {
	es.Lock()
	defer es.Unlock()
	_, ok := es.Employees[id]
	if !ok {
		return fmt.Errorf("employee with ID %d not found", id)
	}
	es.Employees[id] = emp
	return nil
}

// DeleteEmployee deletes an employee from the store by ID
func (es *EmployeeStore) DeleteEmployee(id int) error {
	es.Lock()
	defer es.Unlock()
	_, ok := es.Employees[id]
	if !ok {
		return fmt.Errorf("employee with ID %d not found", id)
	}
	delete(es.Employees, id)
	return nil
}
