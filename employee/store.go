package employee

// PaginationParams holds pagination parameters
type PaginationParams struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

// PaginationResult holds paginated employee records
type PaginationResult struct {
	Total      int        `json:"total"`
	PerPage    int        `json:"per_page"`
	Page       int        `json:"page"`
	TotalPages int        `json:"total_pages"`
	Employees  []Employee `json:"employees"`
}
