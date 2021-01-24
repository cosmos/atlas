package models

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	uuid "github.com/satori/go.uuid"

	"github.com/cosmos/atlas/server/httputil"
)

type (
	// BugTrackerJSON defines the JSON-encodeable type for a ModuleVersion.
	ModuleVersionJSON struct {
		GormModelJSON

		Version       string      `json:"version"`
		Documentation string      `json:"documentation"`
		Repo          string      `json:"repo"`
		SDKCompat     interface{} `json:"sdk_compat"`
		ModuleID      uint        `json:"module_id"`
		PublishedBy   uint        `json:"published_by"`
	}

	// ModuleVersion defines a version associated with a unique module.
	ModuleVersion struct {
		gorm.Model

		Documentation string
		Repo          string `gorm:"not null;default:null"`
		Version       string
		SDKCompat     sql.NullString
		ModuleID      uint
		PublishedBy   uint
	}

	// ModuleKeywords defines the type relationship between a module and all the
	// associated keywords.
	ModuleKeywords struct {
		ModuleID  uint `json:"module_id"`
		KeywordID uint `json:"keyword_id"`
	}

	// ModuleAuthors defines the type relationship between a module and all the
	// associated authors.
	ModuleAuthors struct {
		ModuleID uint `json:"module_id"`
		UserID   uint `json:"user_id"`
	}

	// BugTrackerJSON defines the JSON-encodeable type for a BugTracker.
	BugTrackerJSON struct {
		GormModelJSON

		URL      interface{} `json:"url"`
		Contact  interface{} `json:"contact"`
		ModuleID uint        `json:"module_id"`
	}

	// BugTracker defines the metadata information for reporting bug reports on a
	// given Module type.
	BugTracker struct {
		gorm.Model

		URL      sql.NullString `json:"url"`
		Contact  sql.NullString `json:"contact"`
		ModuleID uint           `json:"module_id"`
	}

	// ModuleJSON defines the JSON-encodeable type for a Module.
	ModuleJSON struct {
		GormModelJSON

		Name        string              `json:"name"`
		Team        string              `json:"team"`
		Description string              `json:"description"`
		Homepage    string              `json:"homepage"`
		Stars       int64               `json:"stars"`
		BugTracker  BugTrackerJSON      `json:"bug_tracker"`
		Keywords    []KeywordJSON       `json:"keywords"`
		Authors     []UserJSON          `json:"authors"`
		Owners      []UserJSON          `json:"owners"`
		Versions    []ModuleVersionJSON `json:"versions"`
	}

	// UserModuleFavorite defines the behavior of a user staring a module record.
	UserModuleFavorite struct {
		UserID   uint `json:"user_id"`
		ModuleID uint `json:"module_id"`
	}

	// ModuleOwnerInvite defines the a module owner invitation relationship.
	ModuleOwnerInvite struct {
		CreatedAt time.Time
		UpdatedAt time.Time

		ModuleID        uint
		InvitedUserID   uint
		InvitedByUserID uint
		Token           uuid.UUID
	}

	// Module defines a Cosmos SDK module.
	Module struct {
		gorm.Model

		Name        string `gorm:"not null;default:null"`
		Team        string `gorm:"not null;default:null"`
		Homepage    string
		Description string
		Stars       int64
		BugTracker  BugTracker      `gorm:"foreignKey:module_id"`
		Keywords    []Keyword       `gorm:"many2many:module_keywords"`
		Authors     []User          `gorm:"many2many:module_authors"`
		Owners      []User          `gorm:"many2many:module_owners"`
		Versions    []ModuleVersion `gorm:"foreignKey:module_id"`

		Version ModuleVersion `gorm:"-"` // current version in manifest
	}
)

// MarshalJSON implements custom JSON marshaling for the ModuleVersion model.
func (mv ModuleVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(mv.NewModuleVersionJSON())
}

func (mv ModuleVersion) NewModuleVersionJSON() ModuleVersionJSON {
	sdkCompat, _ := mv.SDKCompat.Value()

	return ModuleVersionJSON{
		GormModelJSON: GormModelJSON{
			ID:        mv.ID,
			CreatedAt: mv.CreatedAt,
			UpdatedAt: mv.UpdatedAt,
		},
		Documentation: mv.Documentation,
		Repo:          mv.Repo,
		Version:       mv.Version,
		SDKCompat:     sdkCompat,
		ModuleID:      mv.ModuleID,
		PublishedBy:   mv.PublishedBy,
	}
}

// MarshalJSON implements custom JSON marshaling for the BugTracker model.
func (bt BugTracker) MarshalJSON() ([]byte, error) {
	return json.Marshal(bt.NewBugTrackerJSON())
}

func (bt BugTracker) NewBugTrackerJSON() BugTrackerJSON {
	btURL, _ := bt.URL.Value()
	btContact, _ := bt.Contact.Value()

	return BugTrackerJSON{
		GormModelJSON: GormModelJSON{
			ID:        bt.ID,
			CreatedAt: bt.CreatedAt,
			UpdatedAt: bt.UpdatedAt,
		},
		URL:      btURL,
		Contact:  btContact,
		ModuleID: bt.ModuleID,
	}
}

// MarshalJSON implements custom JSON marshaling for the Module model.
func (m Module) MarshalJSON() ([]byte, error) {
	versionsJSON := make([]ModuleVersionJSON, len(m.Versions))
	for i, v := range m.Versions {
		versionsJSON[i] = v.NewModuleVersionJSON()
	}

	ownersJSON := make([]UserJSON, len(m.Owners))
	for i, o := range m.Owners {
		ownersJSON[i] = o.NewUserJSON()
	}

	authorsJSON := make([]UserJSON, len(m.Authors))
	for i, a := range m.Authors {
		authorsJSON[i] = a.NewUserJSON()
	}

	keywordsJSON := make([]KeywordJSON, len(m.Keywords))
	for i, k := range m.Keywords {
		keywordsJSON[i] = k.NewKeywordJSON()
	}

	return json.Marshal(ModuleJSON{
		GormModelJSON: GormModelJSON{
			ID:        m.ID,
			CreatedAt: m.CreatedAt,
			UpdatedAt: m.UpdatedAt,
		},
		Name:        m.Name,
		Team:        m.Team,
		Description: m.Description,
		Homepage:    m.Homepage,
		BugTracker:  m.BugTracker.NewBugTrackerJSON(),
		Keywords:    keywordsJSON,
		Owners:      ownersJSON,
		Authors:     authorsJSON,
		Versions:    versionsJSON,
		Stars:       m.Stars,
	})
}

// BeforeSave implements a GORM hook for updating a Module record before it is
// created or updated.
func (m *Module) BeforeSave(tx *gorm.DB) error {
	var numStars int64

	if err := tx.Model(&UserModuleFavorite{}).Where("module_id = ?", m.ID).Count(&numStars).Error; err != nil {
		return fmt.Errorf("failed to query for module favorites count: %w", err)
	}

	m.Stars = numStars
	m.Name = strings.ToLower(m.Name)
	m.Team = strings.ToLower(m.Team)
	return nil
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

		// append the module version if it is new
		versionQuery := &ModuleVersion{Version: m.Version.Version, ModuleID: record.ID}
		if err := tx.Where(versionQuery).First(&ModuleVersion{}).Error; err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
			modVer := ModuleVersion{
				Documentation: m.Version.Documentation,
				Repo:          m.Version.Repo,
				Version:       m.Version.Version,
				SDKCompat:     m.Version.SDKCompat,
				PublishedBy:   m.Version.PublishedBy,
			}
			if err := tx.Model(&record).Association("Versions").Append(&modVer); err != nil {
				return fmt.Errorf("failed to update module version: %w", err)
			}
		}

		// update primary fields
		if err := tx.First(&record, record.ID).Updates(Module{
			Team:        m.Team,
			Description: m.Description,
			Homepage:    m.Homepage,
			BugTracker:  bugTracker,
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

// Starred returns a boolean defining if a given user by ID has starred a module.
// It returns an error upon query failure.
func (m Module) UserStarred(db *gorm.DB, userID uint) (bool, error) {
	var record UserModuleFavorite

	if err := db.Where(UserModuleFavorite{ModuleID: m.ID, UserID: userID}).First(&record).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}

		return false, fmt.Errorf("failed to query for user module favorite: %w", err)
	}

	return true, nil
}

// Star favorites a module by ID for a given userID. It returns an error upon
// failure, otherwise it returns the total nubmer of favorites for the module.
func (m Module) Star(db *gorm.DB, userID uint) (int64, error) {
	ok, err := m.UserStarred(db, userID)
	if err != nil {
		return 0, err
	}
	if ok {
		// user already favored, so we just return the stars count
		return m.Stars, nil
	}

	if err := db.Create(&UserModuleFavorite{ModuleID: m.ID, UserID: userID}).Error; err != nil {
		return 0, fmt.Errorf("failed to favorite module: %w", err)
	}

	if err := db.Save(&m).Error; err != nil {
		return 0, fmt.Errorf("failed to update module: %w", err)
	}

	return m.Stars, nil
}

// UnStar removes a favorite for a module by ID for a given userID. It returns an
// error upon failure, otherwise it returns the total nubmer of favorites for the
// module.
func (m Module) UnStar(db *gorm.DB, userID uint) (int64, error) {
	ok, err := m.UserStarred(db, userID)
	if err != nil {
		return 0, err
	}
	if !ok {
		// user did not star this module, so we just simply return the existing total
		return m.Stars, nil
	}

	if err := db.Delete(UserModuleFavorite{}, "module_id = ? AND user_id = ?", m.ID, userID).Error; err != nil {
		return 0, fmt.Errorf("failed to remove favorite for module: %w", err)
	}

	if err := db.Save(&m).Error; err != nil {
		return 0, fmt.Errorf("failed to update module: %w", err)
	}

	return m.Stars, nil
}

// QueryModule performs a query for a Module record. The resulting record, if it
// exists, is returned. If the query fails or the record does not exist, an error
// is returned.
func QueryModule(db *gorm.DB, query map[string]interface{}) (Module, error) {
	var record Module

	if err := db.Preload(clause.Associations).Where(query).First(&record).Error; err != nil {
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

// AddOwner adds a given User as an owner to a Module and deletes the corresponding
// ModuleOwnerInvite record. It returns an error upon failure.
func (m Module) AddOwner(db *gorm.DB, owner User) (Module, error) {
	var record Module

	err := db.Transaction(func(tx *gorm.DB) error {
		m.Owners = append(m.Owners, owner)

		module, err := m.Upsert(tx)
		if err != nil {
			return err
		}

		record = module

		if err := tx.Where("module_id = ? AND invited_user_id = ?", m.ID, owner.ID).Delete(ModuleOwnerInvite{}).Error; err != nil {
			return fmt.Errorf("failed to delete module owner invitation: %w", err)
		}

		// commit the tx
		return nil
	})
	if err != nil {
		return Module{}, err
	}

	return record, nil
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

// GetAllModules returns a slice of Module objects paginated by an offset, order
// and limit. An error is returned upon database query failure.
func GetAllModules(db *gorm.DB, pq httputil.PaginationQuery) ([]Module, Paginator, error) {
	var (
		modules []Module
		total   int64
	)

	tx := db.Preload(clause.Associations)

	if err := tx.Scopes(paginateScope(pq, &modules)).Error; err != nil {
		return nil, Paginator{}, fmt.Errorf("failed to query for modules: %w", err)
	}

	if err := db.Model(&Module{}).Count(&total).Error; err != nil {
		return nil, Paginator{}, fmt.Errorf("failed to query for module count: %w", err)
	}

	return modules, buildPaginator(pq, total), nil
}

// SearchModules performs a paginated query for a set of modules by name, team,
// description or keywords. If no matching modules exist, an empty slice is
// returned.
func SearchModules(db *gorm.DB, query string, pq httputil.PaginationQuery) ([]Module, Paginator, error) {
	if len(query) == 0 {
		return []Module{}, Paginator{}, nil
	}

	type queryRow struct {
		ModuleID uint
	}

	// Perform a join on modules and keywords, including modules without keywords,
	// and execute a text search against fields: module name, module team,
	// module description, and keyword. This query returns all the matching module
	// IDs without any pagination or sorting.
	//
	// TODO: Analyze and improve the query as loading all matching IDs in memory
	// is probably not ideal.
	rows, err := db.Raw(`SELECT DISTINCT
  ON (module_id) results.module_id AS module_id
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
  ) AS results;
`, query).Rows()
	if err != nil {
		return nil, Paginator{}, fmt.Errorf("failed to search for modules: %w", err)
	}

	defer rows.Close()

	moduleIDs := []uint{}
	for rows.Next() {
		var qr queryRow
		if err := db.ScanRows(rows, &qr); err != nil {
			return nil, Paginator{}, fmt.Errorf("failed to search for modules: %w", err)
		}

		moduleIDs = append(moduleIDs, qr.ModuleID)
	}

	if len(moduleIDs) == 0 {
		return []Module{}, Paginator{}, nil
	}

	var modules []Module

	if err := db.Preload(clause.Associations).
		Offset(int(offsetFromPage(pq))).
		Limit(int(pq.Limit)).
		Order(buildOrderBy(pq)).
		Find(&modules, moduleIDs).Error; err != nil {
		return nil, Paginator{}, fmt.Errorf("failed to search for modules: %w", err)
	}

	return modules, buildPaginator(pq, int64(len(moduleIDs))), nil
}

// BeforeSave will create and set the ModuleOwnerInvite UUID.
func (moi *ModuleOwnerInvite) BeforeSave(_ *gorm.DB) error {
	moi.Token = uuid.NewV4()
	return nil
}

// Upsert creates or updates a ModuleOwnerInvite record. If no record exists, a
// new record with a UUID token is created. Otherwise, the existing ModuleOwnerInvite
// record's UUID token is updated/regenerated. It returns an error up database
// failure.
func (moi ModuleOwnerInvite) Upsert(db *gorm.DB) (ModuleOwnerInvite, error) {
	var record ModuleOwnerInvite

	query := "module_id = ? AND invited_user_id = ? AND invited_by_user_id = ?"
	err := db.Transaction(func(tx *gorm.DB) error {
		err := tx.Where(query, moi.ModuleID, moi.InvitedUserID, moi.InvitedByUserID).First(&record).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				if err := tx.Create(&moi).Error; err != nil {
					return fmt.Errorf("failed to create module owner invitation: %w", err)
				}

				// commit the tx
				return nil
			} else {
				return fmt.Errorf("failed to query for module owner invitation: %w", err)
			}
		}

		if err := tx.Where(query, moi.ModuleID, moi.InvitedUserID, moi.InvitedByUserID).Save(&record).Error; err != nil {
			return fmt.Errorf("failed to update module owner invitation: %w", err)
		}

		// commit the tx
		return nil
	})
	if err != nil {
		return ModuleOwnerInvite{}, err
	}

	err = db.Where(query, moi.ModuleID, moi.InvitedUserID, moi.InvitedByUserID).First(&record).Error
	return record, err
}

// QueryModuleOwnerInvite performs a query for a ModuleOwnerInvite record.
// The resulting record, if it exists, is returned. If the query fails or the
// record does not exist, an error is returned.
func QueryModuleOwnerInvite(db *gorm.DB, query map[string]interface{}) (ModuleOwnerInvite, error) {
	var record ModuleOwnerInvite

	if err := db.Where(query).First(&record).Error; err != nil {
		return ModuleOwnerInvite{}, fmt.Errorf("failed to query module owner invitation: %w", err)
	}

	return record, nil
}
