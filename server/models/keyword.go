package models

import (
	"fmt"

	"gorm.io/gorm"
)

// Keyword defines a module keyword, where a module can have one or more keywords.
type Keyword struct {
	GormModel

	Name string `json:"name" yaml:"name"`
}

// Query performs a query for a Keyword record where the search criteria is
// defined by the receiver object. The resulting record, if it exists, is
// returned. If the query fails or the record does not exist, an error is returned.
func (k Keyword) Query(db *gorm.DB) (Keyword, error) {
	var record Keyword

	if err := db.Where(k).First(&record).Error; err != nil {
		return Keyword{}, fmt.Errorf("failed to query keyword: %w", err)
	}

	return record, nil
}

// GetAllKeywords returns a slice of Keyword objects paginated by a cursor and a
// limit. The cursor must be the ID of the last retrieved object. An error is
// returned upon database query failure.
func GetAllKeywords(db *gorm.DB, cursor uint, limit int) ([]Keyword, error) {
	var keywords []Keyword

	if err := db.Limit(limit).Order("id asc").Where("id > ?", cursor).Find(&keywords).Error; err != nil {
		return nil, fmt.Errorf("failed to query for keywords: %w", err)
	}

	return keywords, nil
}
