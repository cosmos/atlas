package server

// PaginationResponse defines a generic type encapsulating a paginated response.
// Client should not rely on decoding into this type as the Results is an
// interface.
type PaginationResponse struct {
	Limit   int         `json:"limit" yaml:"limit"`
	Cursor  uint        `json:"cursor" yaml:"cursor"`
	Count   int         `json:"count" yaml:"count"`
	Results interface{} `json:"results" yaml:"results"`
}

func NewPaginationResponse(count, limit int, cursor uint, results interface{}) PaginationResponse {
	return PaginationResponse{
		Limit:   limit,
		Cursor:  cursor,
		Count:   count,
		Results: results,
	}
}
