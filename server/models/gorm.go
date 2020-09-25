package models

import (
	"time"

	"gorm.io/gorm"
)

// GormModel replicates the GormModel type with proper JSON tags defined.
type GormModel struct {
	ID        uint           `gorm:"primarykey" json:"id" yaml:"id"`
	CreatedAt time.Time      `json:"created_at" yaml:"created_at"`
	UpdatedAt time.Time      `json:"updated_at" yaml:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-" yaml:"-"`
}
