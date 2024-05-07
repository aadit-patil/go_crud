package http

import (
	"encoding/json"
	"fmt"
	em "go_crud/employee"
	"net/http"
)

type Handler struct {
	Es *em.EmployeeStore
}

// ListEmployees handles listing employees with pagination
func (h Handler) ListEmployees(w http.ResponseWriter, r *http.Request) {
	// Parse pagination parameters
	params := em.PaginationParams{
		Page:  1,
		Limit: 4,
	}
	es := h.Es

	// Retrieve total number of employees
	totalEmployees := len(es.Employees)
	fmt.Println(es.Employees)
	// Calculate pagination metadata
	totalPages := totalEmployees / params.Limit
	if totalEmployees%params.Limit != 0 {
		totalPages++
	}
	if params.Page < 1 || params.Page > totalPages {
		http.Error(w, "Invalid page number", http.StatusBadRequest)
		return
	}

	// Calculate offset and limit
	offset := (params.Page - 1) * params.Limit
	limit := params.Limit

	// Retrieve employees for the requested page
	es.RLock()
	defer es.RUnlock()
	var employees []em.Employee
	for id := range es.Employees {
		if offset > 0 {
			offset--
			continue
		}
		employees = append(employees, es.Employees[id])
		limit--
		if limit == 0 {
			break
		}
	}

	// Prepare pagination result
	paginationResult := em.PaginationResult{
		Total:      totalEmployees,
		PerPage:    params.Limit,
		Page:       params.Page,
		TotalPages: totalPages,
		Employees:  employees,
	}

	// Convert result to JSON and send response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(paginationResult)
}
