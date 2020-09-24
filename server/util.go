package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func parsePagination(r *http.Request) (uint, int, error) {
	cursorStr := r.URL.Query().Get("cursor")
	cursor, err := strconv.ParseUint(cursorStr, 10, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid cursor param: %w", err)
	}

	limitStr := r.URL.Query().Get("limit")
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
