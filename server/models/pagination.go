package models

import (
	"errors"
	"fmt"
	"strconv"

	"gorm.io/gorm"

	"github.com/cosmos/atlas/server/httputil"
)

// ErrInvalidPaginationQuery defines a sentinel error when an invalid pagination
// query is provided.
var ErrInvalidPaginationQuery = errors.New("invalid pagination query")

// Paginator defines pagination result cursor metadata to determine how to
// make subsequent pagination calls.
type Paginator struct {
	PrevCursor string
	NextCursor string
}

// PrevPageScope builds a scope for executing a previous page query on a table
// by the primary key (id). This scope cannot be used for custom queries.
func PrevPageScope(pq httputil.PaginationQuery, table string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		query := fmt.Sprintf(`SELECT *
    FROM (SELECT * FROM %s WHERE id < ? ORDER BY id DESC LIMIT ?) as prev_page
		ORDER BY id ASC;`, table)
		return db.Raw(query, pq.Cursor, pq.Limit)
	}
}

// NextPageScope builds a scope for executing a next page query on a table
// by the primary key (id). This scope cannot be used for custom queries.
func NextPageScope(pq httputil.PaginationQuery, table string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		query := fmt.Sprintf("SELECT * FROM %s WHERE id > ? ORDER BY id ASC LIMIT ?;", table)
		return db.Raw(query, pq.Cursor, pq.Limit)
	}
}

// BuildPaginator returns a Paginator object with previous and next page cursors
// after completing a pagination query. An error is returned if any query fails.
func BuildPaginator(tx *gorm.DB, pq httputil.PaginationQuery, model interface{}, numRecords int, startID, endID uint) (Paginator, error) {
	paginator := Paginator{}

	if numRecords > 0 {
		var prevCount int64
		if err := tx.Model(model).Where("id < ?", startID).Count(&prevCount).Error; err != nil {
			return Paginator{}, fmt.Errorf("failed to query for previous cursor: %w", err)
		}

		if prevCount > 0 {
			paginator.PrevCursor = strconv.Itoa(int(startID))
		}

		if numRecords == pq.Limit {
			var nextCount int64
			if err := tx.Model(model).Where("id > ?", endID).Count(&nextCount).Error; err != nil {
				return Paginator{}, fmt.Errorf("failed to query for next cursor: %w", err)
			}

			if nextCount > 0 {
				paginator.NextCursor = strconv.Itoa(int(endID))
			}
		}
	}

	return paginator, nil
}
