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

// PaginationResponse defines a generic type encapsulating a paginated response.
// Client should not rely on decoding into this type as the Results is an
// interface.
type PaginationResponse struct {
	Limit   int         `json:"limit"`
	Page    int         `json:"page"`
	Count   int         `json:"count"`
	Results interface{} `json:"results"`
}

func NewPaginationResponse(count, page, limit int, results interface{}) PaginationResponse {
	return PaginationResponse{
		Page:    page,
		Limit:   limit,
		Count:   count,
		Results: results,
	}
}

// ParsePagination parses pagination values (page and limit) from an HTTP
// request returning an error upon failure.
func ParsePagination(req *http.Request) (int, int, error) {
	pageStr := req.URL.Query().Get("page")
	page, err := strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid page param: %w", err)
	}

	if page <= 0 {
		return 0, 0, errors.New("page must be positive")
	}

	limitStr := req.URL.Query().Get("limit")
	limit, err := strconv.ParseInt(limitStr, 10, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid limit param: %w", err)
	}

	return int(page), int(limit), nil
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
