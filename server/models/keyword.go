package models

import (
	"encoding/json"
	"fmt"

	"gorm.io/gorm"
)

type (
	// KeywordJSON defines the JSON-encodeable type for a Keyword.
	KeywordJSON struct {
		GormModelJSON

		Name string `json:"name"`
	}

	// Keyword defines a module keyword, where a module can have one or more keywords.
	Keyword struct {
		gorm.Model

		Name string `json:"name"`
	}
)

// MarshalJSON implements custom JSON marshaling for the Keyword model.
func (k Keyword) MarshalJSON() ([]byte, error) {
	return json.Marshal(KeywordJSON{
		GormModelJSON: GormModelJSON{
			ID:        k.ID,
			CreatedAt: k.CreatedAt,
			UpdatedAt: k.UpdatedAt,
		},
		Name: k.Name,
	})
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

// GetAllKeywords returns a slice of Keyword objects paginated by an offset and a
// limit. An error is returned upon database query failure.
func GetAllKeywords(db *gorm.DB, offset, limit int) ([]Keyword, error) {
	var keywords []Keyword

	if err := db.Limit(limit).Offset(offset).Order("id asc").Find(&keywords).Error; err != nil {
		return nil, fmt.Errorf("failed to query for keywords: %w", err)
	}

	return keywords, nil
}
