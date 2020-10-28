package models

import (
	"errors"
	"fmt"
	"strings"

	"gorm.io/gorm"

	"github.com/cosmos/atlas/server/httputil"
)

// ErrInvalidPaginationQuery defines a sentinel error when an invalid pagination
// query is provided.
var ErrInvalidPaginationQuery = errors.New("invalid pagination query")

// Paginator defines pagination result cursor metadata to determine how to
// make subsequent pagination calls.
type Paginator struct {
	PrevPage int64
	NextPage int64
	Total    int64
}

// paginateScope builds a reusable scope for executing a paginated offset query
// and returning the total count.
func paginateScope(pq httputil.PaginationQuery, dest interface{}) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(int(offsetFromPage(pq))).
			Limit(int(pq.Limit)).
			Order(buildOrderBy(pq)).
			Find(dest)
	}
}

// buildPaginator returns a Paginator object with previous and next page values
// after completing a pagination query.
func buildPaginator(pq httputil.PaginationQuery, total int64) Paginator {
	paginator := Paginator{
		Total: total,
	}

	if hasPrevious(pq) {
		paginator.PrevPage = pq.Page - 1
	}

	if hasNext(pq, total) {
		paginator.NextPage = pq.Page + 1
	}

	return paginator
}

func buildOrderBy(pq httputil.PaginationQuery) string {
	tokens := []string{}
	for _, column := range strings.Split(pq.Order, ",") {
		if pq.Reverse {
			tokens = append(tokens, fmt.Sprintf("%s DESC", column))
		} else {
			tokens = append(tokens, fmt.Sprintf("%s ASC", column))
		}
	}

	return strings.Join(tokens, ",")
}

func offsetFromPage(pq httputil.PaginationQuery) int64 {
	return (pq.Page - 1) * pq.Limit
}

func hasPrevious(pq httputil.PaginationQuery) bool {
	return (offsetFromPage(pq) - pq.Limit) >= 0

}

func hasNext(pq httputil.PaginationQuery, total int64) bool {
	return (offsetFromPage(pq) + pq.Limit) < total
}
