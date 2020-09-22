package models

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type (
	// Keyword defines a module keyword, where a module can have one or more keywords.
	Keyword struct {
		gorm.Model

		Name string `json:"name" yaml:"name"`
	}

	// ModuleVersion defines a version associated with a unique module.
	ModuleVersion struct {
		gorm.Model

		Version  string `json:"version" yaml:"version"`
		ModuleID uint   `json:"-" yaml:"-"`
	}

	// ModuleKeywords defines the type relationship between a module and all the
	// associated keywords.
	ModuleKeywords struct {
		ModuleID  uint
		KeywordID uint
	}

	// ModuleAuthors defines the type relationship between a module and all the
	// associated authors.
	ModuleAuthors struct {
		ModuleID uint
		UserID   uint
	}

	// BugTracker defines the metadata information for reporting bug reports on a
	// given Module type.
	BugTracker struct {
		gorm.Model

		URL      string `gorm:"not null;default:null" json:"url" yaml:"url"`
		Contact  string `gorm:"not null;default:null" json:"contact" yaml:"contact"`
		ModuleID uint
	}

	// Module defines a Cosmos SDK module.

	Module struct {
		gorm.Model

		Name          string `gorm:"not null;default:null" json:"name" yaml:"name"`
		Team          string `gorm:"not null;default:null" json:"team" yaml:"team"`
		Description   string `json:"description" yaml:"description"`
		Documentation string `json:"documentation" yaml:"documentation"`
		Homepage      string `json:"homepage" yaml:"homepage"`
		Repo          string `gorm:"not null;default:null" json:"repo" yaml:"repo"`

		// one-to-one relationships
		BugTracker BugTracker `json:"bug_tracker" yaml:"bug_tracker" gorm:"foreignKey:module_id"`

		// many-to-many relationships
		Keywords []Keyword `gorm:"many2many:module_keywords" json:"keywords" yaml:"keywords"`
		Authors  []User    `gorm:"many2many:module_authors" json:"authors" yaml:"authors"`

		// one-to-many relationships
		Version  string          `gorm:"-" json:"-" yaml:"-"` // current version in manifest
		Versions []ModuleVersion `gorm:"foreignKey:module_id" json:"versions" yaml:"versions"`
	}
)

// Upsert will attempt to either create a new Module record or update an
// existing record. A Module record is considered unique by a (name, team) index.
// In the case of the record existing, all primary and one-to-one fields will be
// updated, where authors and keywords are replaced. If the provided Version
// does not exist, it will be appended to the existing set of version relations.
// An error is returned upon failure. Upon success, the created or updated record
// will be returned.
func (m Module) Upsert(db *gorm.DB) (Module, error) {
	var record Module

	tx := db.Where("name = ? AND team = ?", m.Name, m.Team).First(&record)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		if m.Version == "" {
			return Module{}, errors.New("failed to create module: empty module version")
		}
		if len(m.Authors) == 0 {
			return Module{}, errors.New("failed to create module: empty module authors")
		}

		m.Versions = []ModuleVersion{{Version: m.Version}}

		// record does not exist, so we create it
		if err := db.Create(&m).Error; err != nil {
			return Module{}, fmt.Errorf("failed to create module: %w", err)
		}

		return m, nil
	}

	// record exists, so we update the relevant fields
	tx = db.Preload(clause.Associations).First(&record)

	// retrieve or create all authors and update the association
	for i, u := range m.Authors {
		if err := db.Where(User{Name: u.Name}).FirstOrCreate(&u).Error; err != nil {
			return Module{}, fmt.Errorf("failed to fetch or create author: %w", err)
		}
		m.Authors[i] = u
	}

	if err := db.Model(&record).Association("Authors").Replace(m.Authors); err != nil {
		return Module{}, fmt.Errorf("failed to update module authors: %w", err)
	}

	// retrieve or create all keywords and update the association
	for i, k := range m.Keywords {
		if err := db.Where(Keyword{Name: k.Name}).FirstOrCreate(&k).Error; err != nil {
			return Module{}, fmt.Errorf("failed to fetch or create keyword: %w", err)
		}
		m.Keywords[i] = k
	}

	if err := db.Model(&record).Association("Keywords").Replace(m.Keywords); err != nil {
		return Module{}, fmt.Errorf("failed to update module keywords: %w", err)
	}

	// update the bug tracker association
	if err := db.Model(&record.BugTracker).Updates(m.BugTracker).Error; err != nil {
		return Module{}, fmt.Errorf("failed to update module bug tracker: %w", err)
	}

	// append version if new
	versionQuery := &ModuleVersion{Version: m.Version, ModuleID: record.ID}
	if err := db.Where(versionQuery).First(&ModuleVersion{}).Error; err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		if err := db.Model(&record).Association("Versions").Append(&ModuleVersion{Version: m.Version}); err != nil {
			return Module{}, fmt.Errorf("failed to update module version: %w", err)
		}
	}

	// update primary fields
	if err := tx.Updates(Module{
		Team:          m.Team,
		Description:   m.Description,
		Documentation: m.Documentation,
		Homepage:      m.Homepage,
		Repo:          m.Repo,
	}).Error; err != nil {
		return Module{}, fmt.Errorf("failed to update module: %w", err)
	}

	return record, nil
}

// GetModuleByID returns a module by ID. If the module doesn't exist or if the
// query fails, an error is returned.
func GetModuleByID(db *gorm.DB, id uint) (Module, error) {
	var m Module

	if err := db.First(&m, id).Error; err != nil {
		return Module{}, fmt.Errorf("failed to query for module by ID: %w", err)
	}

	return m, nil
}

// GetAllModules returns a slice of Module objects paginated by a cursor and a
// limit. The cursor must be the ID of the last retrieved object. An error is
// returned upon database query failure.
func GetAllModules(db *gorm.DB, cursor uint, limit int) ([]Module, error) {
	var modules []Module

	if err := db.Limit(limit).Order("id asc").Where("id > ?", cursor).Find(&modules).Error; err != nil {
		return nil, fmt.Errorf("failed to query for modules: %w", err)
	}

	return modules, nil
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
