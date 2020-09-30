package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
)

func parsePagination(req *http.Request) (uint, int, error) {
	cursorStr := req.URL.Query().Get("cursor")
	cursor, err := strconv.ParseUint(cursorStr, 10, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid cursor param: %w", err)
	}

	limitStr := req.URL.Query().Get("limit")
	limit, err := strconv.ParseInt(limitStr, 10, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid limit param: %w", err)
	}

	return uint(cursor), int(limit), nil
}

func respondWithError(w http.ResponseWriter, code int, err error) {
	respondWithJSON(w, code, map[string]string{"error": err.Error()})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, _ = w.Write(response)
}

func transformValidationError(err error) error {
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
