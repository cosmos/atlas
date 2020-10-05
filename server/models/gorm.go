package models

import (
	"database/sql"
	"time"
)

// GormModelJSON defines a wrapper around a gorm.Model object that is used for
// JSON marshaling.
type GormModelJSON struct {
	ID        uint      `json:"id" yaml:"id"`
	CreatedAt time.Time `json:"created_at" yaml:"created_at"`
	UpdatedAt time.Time `json:"updated_at" yaml:"updated_at"`
}

func NewNullString(s string) sql.NullString {
	if s == "" {
		return sql.NullString{}
	}

	return sql.NullString{String: s, Valid: true}
}

func NewNullInt64(i int64) sql.NullInt64 {
	if i == 0 {
		return sql.NullInt64{}
	}

	return sql.NullInt64{Int64: i, Valid: true}
}
