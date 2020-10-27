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

const (
	// PageNext defines an enumerated value for retreiving the next page of
	// paginated data.
	PageNext = "next"

	// PagePrev defines an enumerated value for retreiving the previous page of
	// paginated data.
	PagePrev = "prev"
)

// PaginationQuery defines the structure containing pagination request information
// from client HTTP requests.
type PaginationQuery struct {
	Cursor string
	Page   string
	Limit  int
}

// PaginationResponse defines a generic type encapsulating a paginated response.
// Client should not rely on decoding into this type as the Results is an
// interface.
type PaginationResponse struct {
	Limit      int         `json:"limit"`
	Count      int         `json:"count"`
	PrevCursor string      `json:"prev_cursor"`
	NextCursor string      `json:"next_cursor"`
	Results    interface{} `json:"results"`
}

func NewPaginationResponse(limit, count int, prevC, nextC string, results interface{}) PaginationResponse {
	return PaginationResponse{
		Limit:      limit,
		Count:      count,
		PrevCursor: prevC,
		NextCursor: nextC,
		Results:    results,
	}
}

// ParsePaginationQueryParams parses pagination values from an HTTP request
// returning an error upon failure.
func ParsePaginationQueryParams(req *http.Request) (PaginationQuery, error) {
	cursor := req.URL.Query().Get("cursor")
	if cursor == "" {
		return PaginationQuery{}, errors.New("invalid pagination cursor: cannot be empty")
	}

	cursorInt, err := strconv.ParseInt(cursor, 10, 64)
	if err != nil {
		return PaginationQuery{}, fmt.Errorf("invalid pagination cursor '%s': %w", cursor, err)
	}

	if cursorInt < 0 {
		return PaginationQuery{}, fmt.Errorf("invalid pagination cursor '%s': cursor cannot be negative", cursor)
	}

	limitStr := req.URL.Query().Get("limit")
	limit, err := strconv.ParseInt(limitStr, 10, 64)
	if err != nil {
		return PaginationQuery{}, fmt.Errorf("invalid pagination limit: %w", err)
	}

	if limit < 0 {
		return PaginationQuery{}, fmt.Errorf("invalid pagination limit '%d': limit cannot be negative", limit)
	}

	page := req.URL.Query().Get("page")
	if page != PagePrev && page != PageNext {
		return PaginationQuery{}, fmt.Errorf("invalid pagination page: must be '%s' or '%s'", PagePrev, PageNext)
	}

	return PaginationQuery{
		Cursor: cursor,
		Page:   page,
		Limit:  int(limit),
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
