package httputil

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/InVisionApp/go-health/v2"
	"github.com/InVisionApp/go-health/v2/checkers"
	"github.com/go-playground/validator/v10"
)

// Common HTTP methods and header values
const (
	MethodGET    = "GET"
	MethodPOST   = "POST"
	MethodPUT    = "PUT"
	MethodDELETE = "DELETE"

	BearerSchema = "Bearer "
)

// PaginationQuery defines the structure containing pagination request information
// from client HTTP requests.
type PaginationQuery struct {
	Order   string
	Reverse bool
	Page    int64
	Limit   int64
}

// PaginationResponse defines a generic type encapsulating a paginated response.
// Client should not rely on decoding into this type as the Results is an
// interface.
type PaginationResponse struct {
	Order   string      `json:"order"`
	Reverse bool        `json:"reverse"`
	Page    int64       `json:"page"`
	Limit   int64       `json:"limit"`
	Count   int64       `json:"count"`
	Total   int64       `json:"total"`
	Results interface{} `json:"results"`
}

func NewPaginationResponse(pq PaginationQuery, count, total int64, results interface{}) PaginationResponse {
	return PaginationResponse{
		Order:   pq.Order,
		Reverse: pq.Reverse,
		Page:    pq.Page,
		Limit:   pq.Limit,
		Count:   count,
		Total:   total,
		Results: results,
	}
}

// ParsePaginationQueryParams parses pagination values from an HTTP request
// returning an error upon failure.
func ParsePaginationQueryParams(req *http.Request) (PaginationQuery, error) {
	order := req.URL.Query().Get("order")
	if order == "" {
		order = "id"
	} else {
		addID := true
		tokens := strings.Split(order, ",")

		for _, token := range tokens {
			if strings.ToLower(token) == "id" {
				addID = false
			}
		}

		if addID {
			tokens = append(tokens, "id")
		}

		order = strings.Join(tokens, ",")
	}

	var reverse bool

	if reverseStr := req.URL.Query().Get("reverse"); reverseStr != "" {
		ok, err := strconv.ParseBool(reverseStr)
		if err != nil {
			return PaginationQuery{}, fmt.Errorf("invalid pagination 'reverse' parameter: %w", err)
		}

		reverse = ok
	}

	pageStr := req.URL.Query().Get("page")
	page, err := strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		return PaginationQuery{}, fmt.Errorf("invalid pagination 'page' parameter: %w", err)
	}

	if page < 1 {
		return PaginationQuery{}, errors.New("invalid pagination 'page' parameter: page must be positive")
	}

	limitStr := req.URL.Query().Get("limit")
	limit, err := strconv.ParseInt(limitStr, 10, 64)
	if err != nil {
		return PaginationQuery{}, fmt.Errorf("invalid pagination 'limit' parameter: %w", err)
	}

	if limit < 0 {
		return PaginationQuery{}, errors.New("invalid pagination 'limit' parameter: limit must be non-negative")
	}

	return PaginationQuery{
		Order:   order,
		Reverse: reverse,
		Page:    page,
		Limit:   limit,
	}, nil
}

// ErrResponse defines an HTTP error response.
type ErrResponse struct {
	Error string `json:"error"`
}

// RespondWithError provides an auxiliary function to handle all failed HTTP
// requests.
func RespondWithError(w http.ResponseWriter, code int, err error) {
	RespondWithJSON(w, code, ErrResponse{err.Error()})
}

// RespondWithJSON provides an auxiliary function to return an HTTP response
// with JSON content and an HTTP status code.
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, _ = w.Write(response)
}

// TransformValidationError accepts an error from validation and attempts to
// transform the error message to a more human-readable format.
func TransformValidationError(err error) error {
	if _, ok := err.(*validator.InvalidValidationError); ok {
		return err
	}

	valErrs := err.(validator.ValidationErrors)
	msgs := make([]string, len(valErrs))
	for i, err := range valErrs {
		msgs[i] = fmt.Sprintf("invalid %s: %s validation failed", strings.ToLower(err.Field()), err.Tag())
	}

	return errors.New(strings.Join(msgs, "; "))
}

// CreateHealthChecker returns a health checker instance with all checkers
// registered.
func CreateHealthChecker(db checkers.SQLPinger, disableLog bool) (*health.Health, error) {
	h := health.New()
	if disableLog {
		h.DisableLogging()
	}

	sqlCheck, err := checkers.NewSQL(&checkers.SQLConfig{
		Pinger: db,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create SQL health checker: %w", err)
	}

	if err := h.AddCheck(&health.Config{
		Name:     "psql-check",
		Checker:  sqlCheck,
		Interval: time.Duration(1) * time.Minute,
		Fatal:    true,
	}); err != nil {
		return nil, fmt.Errorf("failed to add health checkers: %w", err)
	}

	if err := h.Start(); err != nil {
		return nil, fmt.Errorf("failed to start health checker: %w", err)
	}

	return h, nil
}
