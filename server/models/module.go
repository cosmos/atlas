package models

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type (
	// ModuleVersion defines a version associated with a unique module.
	ModuleVersion struct {
		gorm.Model

		Version   string         `json:"version" yaml:"version"`
		SDKCompat sql.NullString `json:"sdk_compat" yaml:"sdk_compat"`
		ModuleID  uint           `json:"module_id" yaml:"module_id"`
	}

	// ModuleKeywords defines the type relationship between a module and all the
	// associated keywords.
	ModuleKeywords struct {
		ModuleID  uint `json:"module_id" yaml:"module_id"`
		KeywordID uint `json:"keyword_id" yaml:"keyword_id"`
	}

	// ModuleAuthors defines the type relationship between a module and all the
	// associated authors.
	ModuleAuthors struct {
		ModuleID uint `json:"module_id" yaml:"module_id"`
		UserID   uint `json:"user_id" yaml:"user_id"`
	}

	// BugTracker defines the metadata information for reporting bug reports on a
	// given Module type.
	BugTracker struct {
		gorm.Model

		URL      sql.NullString `json:"url" yaml:"url"`
		Contact  sql.NullString `json:"contact" yaml:"contact"`
		ModuleID uint           `json:"module_id" yaml:"module_id"`
	}

	// Module defines a Cosmos SDK module.

	Module struct {
		gorm.Model

		Name string `gorm:"not null;default:null" json:"name" yaml:"name"`
		Team string `gorm:"not null;default:null" json:"team" yaml:"team"`

		Description   string `json:"description" yaml:"description"`
		Documentation string `json:"documentation" yaml:"documentation"`
		Homepage      string `json:"homepage" yaml:"homepage"`
		Repo          string `gorm:"not null;default:null" json:"repo" yaml:"repo"`

		// one-to-one relationships
		BugTracker BugTracker `json:"bug_tracker" yaml:"bug_tracker" gorm:"foreignKey:module_id"`

		// many-to-many relationships
		Keywords []Keyword `gorm:"many2many:module_keywords" json:"keywords" yaml:"keywords"`
		Authors  []User    `gorm:"many2many:module_authors" json:"authors" yaml:"authors"`
		Owners   []User    `gorm:"many2many:module_owners" json:"owners" yaml:"owners"`

		// one-to-many relationships
		Version  ModuleVersion   `gorm:"-" json:"-" yaml:"-"` // current version in manifest
		Versions []ModuleVersion `gorm:"foreignKey:module_id" json:"versions" yaml:"versions"`
	}
)

// MarshalJSON implements custom JSON marshaling for the ModuleVersion model.
func (mv ModuleVersion) MarshalJSON() ([]byte, error) {
	sdkCompat, _ := mv.SDKCompat.Value()

	return json.Marshal(struct {
		GormModelJSON

		Version   string      `json:"version" yaml:"version"`
		SDKCompat interface{} `json:"sdk_compat" yaml:"sdk_compat"`
		ModuleID  uint        `json:"module_id" yaml:"module_id"`
	}{
		GormModelJSON: GormModelJSON{
			ID:        mv.ID,
			CreatedAt: mv.CreatedAt,
			UpdatedAt: mv.UpdatedAt,
		},
		Version:   mv.Version,
		SDKCompat: sdkCompat,
		ModuleID:  mv.ModuleID,
	})
}

// MarshalJSON implements custom JSON marshaling for the BugTracker model.
func (bt BugTracker) MarshalJSON() ([]byte, error) {
	btURL, _ := bt.URL.Value()
	btContact, _ := bt.Contact.Value()

	return json.Marshal(struct {
		GormModelJSON

		URL      interface{} `json:"url" yaml:"url"`
		Contact  interface{} `json:"contact" yaml:"contact"`
		ModuleID uint        `json:"module_id" yaml:"module_id"`
	}{
		GormModelJSON: GormModelJSON{
			ID:        bt.ID,
			CreatedAt: bt.CreatedAt,
			UpdatedAt: bt.UpdatedAt,
		},
		URL:      btURL,
		Contact:  btContact,
		ModuleID: bt.ModuleID,
	})
}

// MarshalJSON implements custom JSON marshaling for the Module model.
func (m Module) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		GormModelJSON

		Name          string          `json:"name" yaml:"name"`
		Team          string          `json:"team" yaml:"team"`
		Description   string          `json:"description" yaml:"description"`
		Documentation string          `json:"documentation" yaml:"documentation"`
		Homepage      string          `json:"homepage" yaml:"homepage"`
		Repo          string          `json:"repo" yaml:"repo"`
		BugTracker    BugTracker      `json:"bug_tracker" yaml:"bug_tracker"`
		Keywords      []Keyword       `json:"keywords" yaml:"keywords"`
		Authors       []User          `json:"authors" yaml:"authors"`
		Owners        []User          `json:"owners" yaml:"owners"`
		Versions      []ModuleVersion `json:"versions" yaml:"versions"`
	}{
		GormModelJSON: GormModelJSON{
			ID:        m.ID,
			CreatedAt: m.CreatedAt,
			UpdatedAt: m.UpdatedAt,
		},
		Name:          m.Name,
		Team:          m.Team,
		Description:   m.Description,
		Documentation: m.Documentation,
		Homepage:      m.Homepage,
		Repo:          m.Repo,
		BugTracker:    m.BugTracker,
		Keywords:      m.Keywords,
		Authors:       m.Authors,
		Owners:        m.Owners,
		Versions:      m.Versions,
	})
}

// Upsert will attempt to either create a new Module record or update an
// existing record. A Module record is considered unique by a (name, team) index.
// In the case of the record existing, all primary and one-to-one fields will be
// updated, where authors and keywords are replaced. If the provided Version
// does not exist, it will be appended to the existing set of version relations.
// An error is returned upon failure. Upon success, the created or updated record
// will be returned.
func (m Module) Upsert(db *gorm.DB) (Module, error) {
	var record Module

	err := db.Transaction(func(tx *gorm.DB) error {
		// fetch or create owners first before updating the association
		for i, o := range m.Owners {
			if err := tx.Where(User{Name: o.Name}).FirstOrCreate(&m.Owners[i]).Error; err != nil {
				return err
			}
		}

		// fetch or create users first before updating the association
		for i, u := range m.Authors {
			if err := tx.Where(User{Name: u.Name}).FirstOrCreate(&m.Authors[i]).Error; err != nil {
				return err
			}
		}

		// fetch or create keywords first before updating the association
		for i, k := range m.Keywords {
			if err := tx.Where(Keyword{Name: k.Name}).FirstOrCreate(&m.Keywords[i]).Error; err != nil {
				return err
			}
		}

		err := tx.Where("name = ? AND team = ?", m.Name, m.Team).First(&record).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				if m.Version.Version == "" {
					return errors.New("failed to create module: empty module version")
				}
				if len(m.Authors) == 0 {
					return errors.New("failed to create module: empty module authors")
				}

				m.Versions = []ModuleVersion{m.Version}

				// record does not exist, so we create it
				if err := tx.Create(&m).Error; err != nil {
					return fmt.Errorf("failed to create module: %w", err)
				}

				// commit the tx
				return nil
			} else {
				return fmt.Errorf("failed to query module: %w", err)
			}
		}

		// record exists, so we update the relevant fields
		if err := tx.Preload(clause.Associations).First(&record).Error; err != nil {
			return err
		}

		// update owners association
		if err := tx.Model(&record).Association("Owners").Replace(m.Owners); err != nil {
			return fmt.Errorf("failed to update module owners: %w", err)
		}

		// update authors association
		if err := tx.Model(&record).Association("Authors").Replace(m.Authors); err != nil {
			return fmt.Errorf("failed to update module authors: %w", err)
		}

		// update keywords association
		if err := tx.Model(&record).Association("Keywords").Replace(m.Keywords); err != nil {
			return fmt.Errorf("failed to update module keywords: %w", err)
		}

		var bugTracker BugTracker
		if err := tx.Model(&record).Association("BugTracker").Find(&bugTracker); err != nil {
			return err
		}

		bugTracker.ModuleID = record.ID
		bugTracker.URL = m.BugTracker.URL
		bugTracker.Contact = m.BugTracker.Contact
		if err := tx.Save(&bugTracker).Error; err != nil {
			return err
		}

		// append version if new
		versionQuery := &ModuleVersion{Version: m.Version.Version, ModuleID: record.ID}
		if err := tx.Where(versionQuery).First(&ModuleVersion{}).Error; err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
			modVer := ModuleVersion{Version: m.Version.Version, SDKCompat: m.Version.SDKCompat}
			if err := tx.Model(&record).Association("Versions").Append(&modVer); err != nil {
				return fmt.Errorf("failed to update module version: %w", err)
			}
		}

		// update primary fields
		if err := tx.First(&record, record.ID).Updates(Module{
			Team:          m.Team,
			Description:   m.Description,
			Documentation: m.Documentation,
			Homepage:      m.Homepage,
			Repo:          m.Repo,
			BugTracker:    bugTracker,
		}).Error; err != nil {
			return fmt.Errorf("failed to update module: %w", err)
		}

		// commit the tx
		return nil
	})
	if err != nil {
		return Module{}, err
	}

	// fetch & reload associations on the upserted record
	if err := db.Preload(clause.Associations).First(&record).Error; err != nil {
		return Module{}, err
	}

	return record, nil
}

// Query performs a query for a Module record where the search criteria is
// defined by the receiver object. The resulting record, if it exists, is
// returned. If the query fails or the record does not exist, an error is
// returned.
func (m Module) Query(db *gorm.DB) (Module, error) {
	var record Module

	if err := db.Preload(clause.Associations).Where(m).First(&record).Error; err != nil {
		return Module{}, fmt.Errorf("failed to query module: %w", err)
	}

	return record, nil
}

// GetLatestVersion returns a module's latest version record, if the module
// exists.
func (m Module) GetLatestVersion(db *gorm.DB) (ModuleVersion, error) {
	var mv ModuleVersion

	if err := db.Order("created_at desc").Where("module_id = ?", m.ID).First(&mv).Error; err != nil {
		return ModuleVersion{}, fmt.Errorf("failed to get latest module version: %w", err)
	}

	return mv, nil
}

// GetModuleByID returns a module by ID. If the module doesn't exist or if the
// query fails, an error is returned.
func GetModuleByID(db *gorm.DB, id uint) (Module, error) {
	var m Module

	if err := db.Preload(clause.Associations).First(&m, id).Error; err != nil {
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

// SearchModules performs a paginated query for a set of modules by name, team,
// description and set of keywords. If not matching modules exist, an empty slice
// is returned.
func SearchModules(db *gorm.DB, query string, cursor uint, limit int) ([]Module, error) {
	if len(query) == 0 {
		return []Module{}, nil
	}

	type queryRow struct {
		ModuleID   int
		ModuleName string
	}

	// Perform a join on modules and keywords, including modules without keywords,
	// and execute a text search against fields: module name, module team,
	// module description, and keyword.
	rows, err := db.Raw(`SELECT DISTINCT
  ON (module_id) results.module_id AS module_id,
  results.module_name AS module_name
FROM
  (
    SELECT
      m.id AS module_id,
      m.name AS module_name,
      m.team,
      m.description,
      k.name
    FROM
      modules m
      LEFT JOIN
        module_keywords mk
        ON (m.id = mk.module_id)
      LEFT JOIN
        keywords k
        ON (mk.keyword_id = k.id)
    WHERE
      to_tsvector('english', COALESCE(m.name, '') || ' ' || COALESCE(m.team, '') || ' ' || COALESCE(m.description, '') || ' ' || COALESCE(k.name, '')) @@ websearch_to_tsquery('english', ?)
      AND m.id > ?
  )
  AS results
ORDER BY
  module_id LIMIT ?;
`, query, cursor, limit).Rows()
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	moduleIDs := []int{}
	for rows.Next() {
		var qr queryRow
		if err := db.ScanRows(rows, &qr); err != nil {
			return nil, fmt.Errorf("failed to search for modules: %w", err)
		}

		moduleIDs = append(moduleIDs, qr.ModuleID)
	}

	if len(moduleIDs) == 0 {
		return []Module{}, nil
	}

	var modules []Module

	if err := db.Preload(clause.Associations).Order("id asc").Find(&modules, moduleIDs).Error; err != nil {
		return nil, fmt.Errorf("failed to search for modules: %w", err)
	}

	return modules, nil
}
