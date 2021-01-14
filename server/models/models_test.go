package models_test

import (
	"database/sql"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"testing"
	"time"

	"github.com/golang-migrate/migrate/v4"
	migratepg "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"

	"github.com/cosmos/atlas/server/httputil"
	"github.com/cosmos/atlas/server/models"
)

type ModelsTestSuite struct {
	suite.Suite

	m      *migrate.Migrate
	db     *sql.DB
	gormDB *gorm.DB
}

// SetupSuite executes once before the suite's tests are executed.
func (mts *ModelsTestSuite) SetupSuite() {
	migrationsPath := os.Getenv("ATLAS_MIGRATIONS_DIR")
	mts.Require().NotEmpty(migrationsPath)

	connStr := os.Getenv("ATLAS_TEST_DATABASE_URL")
	mts.Require().NotEmpty(connStr)

	db, err := sql.Open("postgres", connStr)
	mts.Require().NoError(err)
	mts.Require().NoError(db.Ping())

	driver, err := migratepg.WithInstance(db, &migratepg.Config{})
	mts.Require().NoError(err)

	m, err := migrate.NewWithDatabaseInstance(fmt.Sprintf("file:///%s", migrationsPath), "postgres", driver)
	mts.Require().NoError(err)

	gormDB, err := gorm.Open(
		postgres.Open(connStr),
		&gorm.Config{
			Logger: gormlogger.Discard,
			NowFunc: func() time.Time {
				// Ignore the microseconds so we can use reflect.DeepEqual as the database
				// does not store the same resolution.
				return time.Now().Local().Truncate(time.Microsecond)
			},
		},
	)
	mts.Require().NoError(err)

	mts.m = m
	mts.db = db
	mts.gormDB = gormDB
}

// TearDownSuite executes after all the suite's test have finished.
func (mts *ModelsTestSuite) TearDownSuite() {
	mts.Require().NoError(mts.db.Close())
}

func TestModelsTestSuite(t *testing.T) {
	suite.Run(t, new(ModelsTestSuite))
}

func (mts *ModelsTestSuite) TestModuleCreate() {
	mts.resetDB()

	testCases := []struct {
		name      string
		module    models.Module
		expectErr bool
	}{
		{
			name:      "create module invalid name",
			module:    models.Module{},
			expectErr: true,
		},
		{
			name:      "create module invalid team",
			module:    models.Module{Name: "x/bank"},
			expectErr: true,
		},
		{
			name:      "create module no repo",
			module:    models.Module{Name: "x/bank", Team: "cosmonauts"},
			expectErr: true,
		},
		{
			name: "create module no authors",
			module: models.Module{
				Name: "x/bank",
				Team: "cosmonauts",
			},
			expectErr: true,
		},
		{
			name: "create module no version",
			module: models.Module{
				Name: "x/bank",
				Team: "cosmonauts",
				Authors: []models.User{
					{Name: "foo", Email: models.NewNullString("foo@email.com")},
				},
			},
			expectErr: true,
		},
		{
			name: "create module",
			module: models.Module{
				Name: "x/bank",
				Team: "cosmonauts",
				Authors: []models.User{
					{Name: "foo", Email: models.NewNullString("foo@cosmonauts.com")},
				},
				Version: models.ModuleVersion{
					Version:       "v1.0.0",
					Repo:          "https://github.com/cosmos/cosmos-sdk/releases/tag/v0.39.1",
					Documentation: "https://raw.githubusercontent.com/cosmos/cosmos-sdk/v0.39.1/x/bank/README.md",
				},
				Keywords: []models.Keyword{
					{Name: "tokens"},
				},
				BugTracker: models.BugTracker{
					URL:     models.NewNullString("cosmonauts.com"),
					Contact: models.NewNullString("contact@cosmonauts.com"),
				},
			},
			expectErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc

		mts.Run(tc.name, func() {
			record, err := tc.module.Upsert(mts.gormDB)
			if tc.expectErr {
				mts.Require().Error(err)
			} else {
				mts.Require().NoError(err)
				mts.Require().Equal(tc.module.Name, record.Name)
				mts.Require().Equal(tc.module.Team, record.Team)
				mts.Require().Equal(tc.module.Description, record.Description)
				mts.Require().Equal(tc.module.Homepage, record.Homepage)
				mts.Require().Equal("https://raw.githubusercontent.com/cosmos/cosmos-sdk/v0.39.1/x/bank/README.md", record.Versions[0].Documentation)
				mts.Require().Equal("https://github.com/cosmos/cosmos-sdk/releases/tag/v0.39.1", record.Versions[0].Repo)
			}
		})
	}
}

func (mts *ModelsTestSuite) TestModuleUpdateBasic() {
	mts.resetDB()

	mod := models.Module{
		Name:        "x/bank",
		Team:        "cosmonauts",
		Description: "test description",
		Homepage:    "https://old.cosmos.network",
		Authors: []models.User{
			{Name: "admin"},
		},
		Version: models.ModuleVersion{
			Version:       "v1.0.0",
			Repo:          "https://github.com/cosmos/cosmos-sdk/releases/tag/v0.39.1",
			Documentation: "https://raw.githubusercontent.com/cosmos/cosmos-sdk/v0.39.1/x/bank/README.md",
		},
		Keywords: []models.Keyword{
			{Name: "tokens"},
		},
		BugTracker: models.BugTracker{},
	}

	record, err := mod.Upsert(mts.gormDB)
	mts.Require().NoError(err)

	mod = models.Module{
		Name:        record.Name,
		Team:        record.Team,
		Description: "new test description",
		Homepage:    "https://new.cosmos.network",
		Keywords: []models.Keyword{
			{Name: "tokens"},
			{Name: "module"},
		},
	}

	record, err = mod.Upsert(mts.gormDB)
	mts.Require().NoError(err)
	mts.Require().Equal(record.Team, mod.Team)
	mts.Require().Equal(record.Name, mod.Name)
	mts.Require().Equal(record.Description, mod.Description)
	mts.Require().Equal(record.Homepage, mod.Homepage)
	mts.Require().Len(record.Keywords, 2)
	mts.Require().Equal("tokens", record.Keywords[0].Name)
	mts.Require().Equal("module", record.Keywords[1].Name)
	mts.Require().Equal("https://raw.githubusercontent.com/cosmos/cosmos-sdk/v0.39.1/x/bank/README.md", record.Versions[0].Documentation)
	mts.Require().Equal("https://github.com/cosmos/cosmos-sdk/releases/tag/v0.39.1", record.Versions[0].Repo)
}

func (mts *ModelsTestSuite) TestGetModuleByID() {
	mts.resetDB()

	mod := models.Module{
		Name: "x/bank",
		Team: "cosmonauts",
		Authors: []models.User{
			{Name: "foo", Email: models.NewNullString("foo@cosmonauts.com")},
		},
		Version: models.ModuleVersion{
			Version:       "v1.0.0",
			Repo:          "https://github.com/cosmos/cosmos-sdk/releases/tag/v0.39.1",
			Documentation: "https://raw.githubusercontent.com/cosmos/cosmos-sdk/v0.39.1/x/bank/README.md",
		},
		Keywords: []models.Keyword{
			{Name: "tokens"},
		},
		BugTracker: models.BugTracker{
			URL:     models.NewNullString("cosmonauts.com"),
			Contact: models.NewNullString("contact@cosmonauts.com"),
		},
	}

	mod, err := mod.Upsert(mts.gormDB)
	mts.Require().NoError(err)

	mts.Run("no module exists", func() {
		result, err := models.GetModuleByID(mts.gormDB, mod.ID+1)
		mts.Require().Error(err)
		mts.Require().Equal(models.Module{}, result)
	})

	mts.Run("module exists", func() {
		result, err := models.GetModuleByID(mts.gormDB, mod.ID)
		mts.Require().NoError(err)
		mts.Require().Equal(mod.Name, result.Name)
		mts.Require().Equal(mod.Team, result.Team)
		mts.Require().Equal(mod.Description, result.Description)
		mts.Require().Equal(mod.Homepage, result.Homepage)
		mts.Require().Equal(mod.Versions[0].Documentation, result.Versions[0].Documentation)
		mts.Require().Equal(mod.Versions[0].Repo, result.Versions[0].Repo)
	})
}

func (mts *ModelsTestSuite) TestGetAllModules() {
	mts.resetDB()

	mods, paginator, err := models.GetAllModules(mts.gormDB, httputil.PaginationQuery{Page: 1, Limit: 10, Order: "id"})
	mts.Require().NoError(err)
	mts.Require().Empty(mods)
	mts.Require().Zero(paginator.PrevPage)
	mts.Require().Equal(int64(0), paginator.NextPage)

	for i := 0; i < 25; i++ {
		mod := models.Module{
			Name: fmt.Sprintf("x/bank-%d", i),
			Team: "cosmonauts",
			Authors: []models.User{
				{Name: "foo", Email: models.NewNullString("foo@cosmonauts.com")},
			},
			Version: models.ModuleVersion{
				Version:       "v1.0.0",
				Repo:          "https://github.com/cosmos/cosmos-sdk/releases/tag/v0.39.1",
				Documentation: fmt.Sprintf("https://raw.githubusercontent.com/cosmos/cosmos-sdk/v0.39.1/x/bank-%d/README.md", i),
			},
			Keywords: []models.Keyword{
				{Name: "tokens"},
			},
			BugTracker: models.BugTracker{
				URL:     models.NewNullString("cosmonauts.com"),
				Contact: models.NewNullString("contact@cosmonauts.com"),
			},
		}

		_, err := mod.Upsert(mts.gormDB)
		mts.Require().NoError(err)
	}

	// first page (full) ordered by newest
	mods, paginator, err = models.GetAllModules(mts.gormDB, httputil.PaginationQuery{Page: 1, Limit: 10, Order: "created_at,id", Reverse: true})
	mts.Require().NoError(err)
	mts.Require().Len(mods, 10)
	mts.Require().Zero(paginator.PrevPage)
	mts.Require().Equal(int64(2), paginator.NextPage)
	mts.Require().Equal(uint(25), mods[0].ID)
	mts.Require().Equal(uint(16), mods[len(mods)-1].ID)

	// first page (full)
	mods, paginator, err = models.GetAllModules(mts.gormDB, httputil.PaginationQuery{Page: 1, Limit: 10, Order: "id"})
	mts.Require().NoError(err)
	mts.Require().Len(mods, 10)
	mts.Require().Zero(paginator.PrevPage)
	mts.Require().Equal(int64(2), paginator.NextPage)

	// second page (full)
	mods, paginator, err = models.GetAllModules(mts.gormDB, httputil.PaginationQuery{Page: 2, Limit: 10, Order: "id"})
	mts.Require().NoError(err)
	mts.Require().Len(mods, 10)
	mts.Require().Equal(int64(1), paginator.PrevPage)
	mts.Require().Equal(int64(3), paginator.NextPage)

	// third page (partially full)
	mods, paginator, err = models.GetAllModules(mts.gormDB, httputil.PaginationQuery{Page: 3, Limit: 10, Order: "id"})
	mts.Require().NoError(err)
	mts.Require().Len(mods, 5)
	mts.Require().Equal(int64(2), paginator.PrevPage)
	mts.Require().Zero(paginator.NextPage)
}

func (mts *ModelsTestSuite) TestModuleUpdateBugTracker() {
	mts.resetDB()

	mod := models.Module{
		Name: "x/bank",
		Team: "cosmonauts",
		Authors: []models.User{
			{Name: "admin"},
		},
		Version: models.ModuleVersion{
			Version:       "v1.0.0",
			Repo:          "https://github.com/cosmos/cosmos-sdk/releases/tag/v0.39.1",
			Documentation: "https://raw.githubusercontent.com/cosmos/cosmos-sdk/v0.39.1/x/bank/README.md",
		},
		BugTracker: models.BugTracker{},
	}

	record, err := mod.Upsert(mts.gormDB)
	mts.Require().NoError(err)

	mod.BugTracker = models.BugTracker{
		URL:     models.NewNullString("https://cosmos.network/bugs"),
		Contact: models.NewNullString("bugs@cosmos.network"),
	}

	record, err = mod.Upsert(mts.gormDB)
	mts.Require().NoError(err)
	mts.Require().Equal(mod.BugTracker.URL, record.BugTracker.URL)
	mts.Require().Equal(mod.BugTracker.Contact, record.BugTracker.Contact)

	mod.BugTracker = models.BugTracker{}

	record, err = mod.Upsert(mts.gormDB)
	mts.Require().NoError(err)
	mts.Require().Equal(mod.BugTracker.URL, record.BugTracker.URL)
	mts.Require().Equal(mod.BugTracker.Contact, record.BugTracker.Contact)
}

func (mts *ModelsTestSuite) TestModuleUpdateKeywords() {
	mts.resetDB()

	mod := models.Module{
		Name: "x/bank",
		Team: "cosmonauts",
		Authors: []models.User{
			{Name: "admin"},
		},
		Version: models.ModuleVersion{
			Version:       "v1.0.0",
			Repo:          "https://github.com/cosmos/cosmos-sdk/releases/tag/v0.39.1",
			Documentation: "https://raw.githubusercontent.com/cosmos/cosmos-sdk/v0.39.1/x/bank/README.md",
		},
		BugTracker: models.BugTracker{},
	}

	record, err := mod.Upsert(mts.gormDB)
	mts.Require().NoError(err)

	mod.Keywords = []models.Keyword{
		{Name: "keyword1"}, {Name: "keyword2"}, {Name: "keyword3"},
	}

	record, err = mod.Upsert(mts.gormDB)
	mts.Require().NoError(err)
	mts.Require().Equal(mod.Keywords, record.Keywords)

	mod.Keywords = []models.Keyword{
		{Name: "keyword1"}, {Name: "keyword3"},
	}

	record, err = mod.Upsert(mts.gormDB)
	mts.Require().NoError(err)
	mts.Require().Equal(mod.Keywords, record.Keywords)

	mod.Keywords = []models.Keyword{}

	record, err = mod.Upsert(mts.gormDB)
	mts.Require().NoError(err)
	mts.Require().Equal(mod.Keywords, record.Keywords)
}

func (mts *ModelsTestSuite) TestModuleUpdateAuthors() {
	mts.resetDB()

	mod := models.Module{
		Name: "x/bank",
		Team: "cosmonauts",
		Authors: []models.User{
			{Name: "admin"},
		},
		Owners: []models.User{
			{Name: "admin"},
		},
		Version: models.ModuleVersion{
			Version:       "v1.0.0",
			Repo:          "https://github.com/cosmos/cosmos-sdk/releases/tag/v0.39.1",
			Documentation: "https://raw.githubusercontent.com/cosmos/cosmos-sdk/v0.39.1/x/bank/README.md",
		},
		BugTracker: models.BugTracker{},
	}

	record, err := mod.Upsert(mts.gormDB)
	mts.Require().NoError(err)

	mod.Authors = []models.User{
		{Name: "admin"}, {Name: "user1"}, {Name: "user2"},
	}

	record, err = mod.Upsert(mts.gormDB)
	mts.Require().NoError(err)

	sort.Slice(mod.Authors, func(i, j int) bool { return mod.Authors[i].ID < mod.Authors[j].ID })
	sort.Slice(record.Authors, func(i, j int) bool { return record.Authors[i].ID < record.Authors[j].ID })
	for i := 0; i < len(record.Authors); i++ {
		mts.Require().True(mod.Authors[i].Equal(record.Authors[i]))
	}
	for i := 0; i < len(record.Owners); i++ {
		mts.Require().True(mod.Owners[i].Equal(record.Owners[i]))
	}

	mod.Authors = []models.User{
		{Name: "admin"}, {Name: "user2"},
	}

	record, err = mod.Upsert(mts.gormDB)
	mts.Require().NoError(err)

	sort.Slice(mod.Authors, func(i, j int) bool { return mod.Authors[i].ID < mod.Authors[j].ID })
	sort.Slice(record.Authors, func(i, j int) bool { return record.Authors[i].ID < record.Authors[j].ID })
	for i := 0; i < len(record.Authors); i++ {
		mts.Require().True(mod.Authors[i].Equal(record.Authors[i]))
	}
	for i := 0; i < len(record.Owners); i++ {
		mts.Require().True(mod.Owners[i].Equal(record.Owners[i]))
	}

	mod.Authors = []models.User{}

	record, err = mod.Upsert(mts.gormDB)
	mts.Require().NoError(err)
	mts.Require().Equal(mod.Authors, record.Authors)
	mts.Require().Equal(mod.Owners, record.Owners)
}

func (mts *ModelsTestSuite) TestModuleUpdateOwners() {
	mts.resetDB()

	mod := models.Module{
		Name: "x/bank",
		Team: "cosmonauts",
		Authors: []models.User{
			{Name: "admin"},
		},
		Owners: []models.User{
			{Name: "admin"},
		},
		Version: models.ModuleVersion{
			Version:       "v1.0.0",
			Repo:          "https://github.com/cosmos/cosmos-sdk/releases/tag/v0.39.1",
			Documentation: "https://raw.githubusercontent.com/cosmos/cosmos-sdk/v0.39.1/x/bank/README.md",
		},
		BugTracker: models.BugTracker{},
	}

	record, err := mod.Upsert(mts.gormDB)
	mts.Require().NoError(err)

	mod.Owners = []models.User{
		{Name: "admin"}, {Name: "user1"}, {Name: "user2"},
	}

	record, err = mod.Upsert(mts.gormDB)
	mts.Require().NoError(err)
	mts.Require().Equal(mod.Authors, record.Authors)

	sort.Slice(mod.Owners, func(i, j int) bool { return mod.Owners[i].ID < mod.Owners[j].ID })
	sort.Slice(record.Owners, func(i, j int) bool { return record.Owners[i].ID < record.Owners[j].ID })
	for i := 0; i < len(record.Authors); i++ {
		mts.Require().True(mod.Authors[i].Equal(record.Authors[i]))
	}
	for i := 0; i < len(record.Owners); i++ {
		mts.Require().True(mod.Owners[i].Equal(record.Owners[i]))
	}

	mod.Owners = []models.User{
		{Name: "admin"}, {Name: "user2"},
	}

	record, err = mod.Upsert(mts.gormDB)
	mts.Require().NoError(err)
	mts.Require().Equal(mod.Authors, record.Authors)

	sort.Slice(mod.Owners, func(i, j int) bool { return mod.Owners[i].ID < mod.Owners[j].ID })
	sort.Slice(record.Owners, func(i, j int) bool { return record.Owners[i].ID < record.Owners[j].ID })
	for i := 0; i < len(record.Authors); i++ {
		mts.Require().True(mod.Authors[i].Equal(record.Authors[i]))
	}
	for i := 0; i < len(record.Owners); i++ {
		mts.Require().True(mod.Owners[i].Equal(record.Owners[i]))
	}

	mod.Owners = []models.User{}

	record, err = mod.Upsert(mts.gormDB)
	mts.Require().NoError(err)
	mts.Require().Equal(mod.Authors, record.Authors)
	mts.Require().Equal(mod.Owners, record.Owners)
}

func (mts *ModelsTestSuite) TestModuleUpdateVersion() {
	mts.resetDB()

	mod := models.Module{
		Name: "x/bank",
		Team: "cosmonauts",
		Authors: []models.User{
			{Name: "admin"},
		},
		Owners: []models.User{
			{Name: "admin"},
		},
		Version: models.ModuleVersion{
			Version:       "v1.0.0",
			Repo:          "https://github.com/cosmos/cosmos-sdk/releases/tag/v0.39.1",
			Documentation: "https://raw.githubusercontent.com/cosmos/cosmos-sdk/v0.39.1/x/bank/README.md",
		},
		BugTracker: models.BugTracker{},
	}

	record, err := mod.Upsert(mts.gormDB)
	mts.Require().NoError(err)

	// update version
	mod.Version = models.ModuleVersion{
		Version:       "v1.0.1",
		Repo:          "https://github.com/cosmos/cosmos-sdk/releases/tag/v0.39.2",
		Documentation: "https://raw.githubusercontent.com/cosmos/cosmos-sdk/v0.39.2/x/bank/README.md",
	}

	record, err = mod.Upsert(mts.gormDB)
	mts.Require().NoError(err)
	mts.Require().Len(record.Versions, 2)

	latest, err := record.GetLatestVersion(mts.gormDB)
	mts.Require().NoError(err)
	mts.Require().Equal(mod.Version.Version, latest.Version)
	mts.Require().Equal(mod.Version.SDKCompat, latest.SDKCompat)

	// no version update
	mod.Version = models.ModuleVersion{
		Version:       "v1.0.1",
		Repo:          "https://github.com/cosmos/cosmos-sdk/releases/tag/v0.39.2",
		Documentation: "https://raw.githubusercontent.com/cosmos/cosmos-sdk/v0.39.2/x/bank/README.md",
	}

	record, err = mod.Upsert(mts.gormDB)
	mts.Require().NoError(err)
	mts.Require().Len(record.Versions, 2)

	latest, err = record.GetLatestVersion(mts.gormDB)
	mts.Require().NoError(err)
	mts.Require().Equal(mod.Version.Version, latest.Version)
	mts.Require().Equal(mod.Version.SDKCompat, latest.SDKCompat)

	// update version again
	mod.Version = models.ModuleVersion{
		Version:       "v2.0.0",
		Repo:          "https://github.com/cosmos/cosmos-sdk/releases/tag/v0.40.0",
		Documentation: "https://raw.githubusercontent.com/cosmos/cosmos-sdk/v0.40.0/x/bank/README.md",
	}

	record, err = mod.Upsert(mts.gormDB)
	mts.Require().NoError(err)
	mts.Require().Len(record.Versions, 3)

	latest, err = record.GetLatestVersion(mts.gormDB)
	mts.Require().NoError(err)
	mts.Require().Equal(mod.Version.Version, latest.Version)
	mts.Require().Equal(mod.Version.SDKCompat, latest.SDKCompat)
}

func (mts *ModelsTestSuite) TestModuleSearch() {
	mts.resetDB()

	teams := []string{"teamA", "teamB", "teamC", "teamD"}
	bugTrackers := []models.BugTracker{
		{URL: models.NewNullString("teamA.com"), Contact: models.NewNullString("contact@teamA.com")},
		{URL: models.NewNullString("teamB.com"), Contact: models.NewNullString("contact@teamB.com")},
		{URL: models.NewNullString("teamC.com"), Contact: models.NewNullString("contact@teamC.com")},
		{URL: models.NewNullString("teamD.com"), Contact: models.NewNullString("contact@teamD.com")},
	}

	godUser := models.User{Name: "deus", Email: models.NewNullString("deus@email.com")}
	users := []models.User{
		{Name: "userA", Email: models.NewNullString("usera@email.com")},
		{Name: "userB", Email: models.NewNullString("userb@email.com")},
		{Name: "userC", Email: models.NewNullString("userc@email.com")},
		{Name: "userD", Email: models.NewNullString("userd@email.com")},
	}

	mods, paginator, err := models.SearchModules(mts.gormDB, "test", httputil.PaginationQuery{Page: 1, Limit: 10, Order: "id"})
	mts.Require().NoError(err)
	mts.Require().Empty(mods)
	mts.Require().Zero(paginator.PrevPage)
	mts.Require().Zero(paginator.NextPage)

	for i := 0; i < 10; i++ {
		randomIndex := rand.Intn(len(teams))
		randTeam := teams[randomIndex]
		randBugTracker := bugTrackers[randomIndex]

		randomIndex = rand.Intn(len(users))
		randUser := users[randomIndex]

		mod := models.Module{
			Name: fmt.Sprintf("x/mod-%d", i),
			Team: randTeam,
			Authors: []models.User{
				godUser,
				randUser,
			},
			Version: models.ModuleVersion{
				Version:       fmt.Sprintf("v1.0.%d", i),
				Repo:          "https://github.com/cosmos/cosmos-sdk/releases/tag/v0.39.1",
				Documentation: fmt.Sprintf("https://raw.githubusercontent.com/cosmos/cosmos-sdk/v0.39.1/x/mod-%d/README.md", i),
			},
			Keywords: []models.Keyword{
				{Name: "module"},
				{Name: fmt.Sprintf("mod-keyword-%d", i+1)},
				{Name: fmt.Sprintf("mod-keyword-%d", (i+1)*3)},
			},
			BugTracker: randBugTracker,
		}

		_, err := mod.Upsert(mts.gormDB)
		mts.Require().NoError(err)
	}

	testCases := []struct {
		name              string
		query             string
		pageQuery         httputil.PaginationQuery
		expectedRecords   map[string]bool
		expectedPaginator models.Paginator
	}{
		{
			"empty query",
			"",
			httputil.PaginationQuery{Page: 1, Limit: 5, Order: "id"},
			map[string]bool{},
			models.Paginator{PrevPage: 0, NextPage: 0, Total: 0},
		},
		{
			"no matching query",
			"no match",
			httputil.PaginationQuery{Page: 1, Limit: 5, Order: "id"},
			map[string]bool{},
			models.Paginator{PrevPage: 0, NextPage: 0, Total: 0},
		},
		{
			"matches one record",
			"x/mod-1",
			httputil.PaginationQuery{Page: 1, Limit: 5, Order: "id"},
			map[string]bool{"x/mod-1": true},
			models.Paginator{PrevPage: 0, NextPage: 0, Total: 1},
		},
		{
			"matches all records (page 1)",
			"module",
			httputil.PaginationQuery{Page: 1, Limit: 5, Order: "id"},
			map[string]bool{
				"x/mod-0": true,
				"x/mod-1": true,
				"x/mod-2": true,
				"x/mod-3": true,
				"x/mod-4": true,
			},
			models.Paginator{PrevPage: 0, NextPage: 2, Total: 10},
		},
		{
			"matches all records (page 2)",
			"module",
			httputil.PaginationQuery{Page: 2, Limit: 5, Order: "id"},
			map[string]bool{
				"x/mod-5": true,
				"x/mod-6": true,
				"x/mod-7": true,
				"x/mod-8": true,
				"x/mod-9": true,
			},
			models.Paginator{PrevPage: 1, NextPage: 0, Total: 10},
		},
	}

	for _, tc := range testCases {
		tc := tc

		mts.Run(tc.name, func() {
			modules, paginator, err := models.SearchModules(mts.gormDB, tc.query, tc.pageQuery)
			mts.Require().NoError(err)
			mts.Require().Len(modules, len(tc.expectedRecords))
			mts.Require().Equal(tc.expectedPaginator, paginator)

			for _, m := range modules {
				mts.Require().Contains(tc.expectedRecords, m.Name)
			}
		})
	}
}

func (mts *ModelsTestSuite) TestUserTokens() {
	mts.resetDB()

	u := models.User{
		Name:              "foo",
		GithubUserID:      models.NewNullInt64(12345),
		GithubAccessToken: models.NewNullString("access_token"),
		Email:             models.NewNullString("foo@email.com"),
		AvatarURL:         "https://avatars.com/myavatar.jpg",
		GravatarID:        "gravatar_id",
	}

	record, err := u.Upsert(mts.gormDB)
	mts.Require().NoError(err)
	mts.Require().Equal(int64(0), record.CountTokens(mts.gormDB))

	token1, err := record.CreateToken(mts.gormDB, "dev")
	mts.Require().NoError(err)
	mts.Require().NotEmpty(token1.Token)

	mts.Require().Equal(uint(0), token1.Count)
	for i := 0; i < 25; i++ {
		token1, err = token1.IncrCount(mts.gormDB)
		mts.Require().NoError(err)
		mts.Require().Equal(uint(i+1), token1.Count)
	}

	token2, err := record.CreateToken(mts.gormDB, "prod")
	mts.Require().NoError(err)
	mts.Require().NotEmpty(token2.Token)

	mts.Require().Equal(int64(2), record.CountTokens(mts.gormDB))
	mts.Require().NotEqual(token1, token2)

	tokens, err := record.GetTokens(mts.gormDB)
	mts.Require().NoError(err)
	mts.Require().Len(tokens, 2)

	token2, err = token2.Revoke(mts.gormDB)
	mts.Require().NoError(err)
	mts.Require().True(token2.Revoked)

	token, err := models.QueryUserToken(mts.gormDB, map[string]interface{}{"token": token1.Token.String()})
	mts.Require().NoError(err)
	mts.Require().Equal(token1, token)

	// duplicate name
	token3, err := record.CreateToken(mts.gormDB, "prod")
	mts.Require().Error(err)
	mts.Require().Equal(uuid.UUID{}, token3.Token)
}

func (mts *ModelsTestSuite) TestUserUpsert() {
	mts.resetDB()

	testCases := []struct {
		name      string
		user      models.User
		expectErr bool
	}{
		{
			"valid user",
			models.User{
				Name:              "foo",
				GithubUserID:      models.NewNullInt64(12345),
				GithubAccessToken: models.NewNullString("access_token"),
				Email:             models.NewNullString("foo@email.com"),
				AvatarURL:         "https://avatars.com/myavatar.jpg",
				GravatarID:        "gravatar_id",
			},
			false,
		},
		{
			"updated user github id",
			models.User{
				Name:              "foo",
				GithubUserID:      models.NewNullInt64(67899),
				GithubAccessToken: models.NewNullString("access_token"),
				Email:             models.NewNullString("foo@email.com"),
				AvatarURL:         "https://avatars.com/myavatar.jpg",
				GravatarID:        "gravatar_id",
			},
			false,
		},
		{
			"updated user email",
			models.User{
				Name:              "foo",
				GithubUserID:      models.NewNullInt64(12345),
				GithubAccessToken: models.NewNullString("access_token"),
				Email:             models.NewNullString("newfoo@email.com"),
				AvatarURL:         "https://avatars.com/myavatar.jpg",
				GravatarID:        "gravatar_id",
			},
			false,
		},
		{
			"updated user avatar url",
			models.User{
				Name:              "foo",
				GithubUserID:      models.NewNullInt64(12345),
				GithubAccessToken: models.NewNullString("access_token"),
				Email:             models.NewNullString("foo@email.com"),
				AvatarURL:         "https://avatars.com/mynewavatar.jpg",
				GravatarID:        "gravatar_id",
			},
			false,
		},
		{
			"updated user gravatar id",
			models.User{
				Name:              "foo",
				GithubUserID:      models.NewNullInt64(12345),
				GithubAccessToken: models.NewNullString("access_token"),
				Email:             models.NewNullString("foo@email.com"),
				AvatarURL:         "https://avatars.com/myavatar.jpg",
				GravatarID:        "new_gravatar_id",
			},
			false,
		},
		{
			"second valid user",
			models.User{
				Name:              "bar",
				GithubUserID:      models.NewNullInt64(567899),
				GithubAccessToken: models.NewNullString("access_token_bar"),
				Email:             models.NewNullString("bar@email.com"),
				AvatarURL:         "https://avatars.com/baravatar.jpg",
				GravatarID:        "bar_gravatar_id",
			},
			false,
		},
	}

	for _, tc := range testCases {
		tc := tc

		mts.Run(tc.name, func() {
			record, err := tc.user.Upsert(mts.gormDB)
			if tc.expectErr {
				mts.Require().Error(err)
			} else {
				mts.Require().NoError(err)
				mts.Require().Equal(tc.user.Name, record.Name)
				mts.Require().Equal(tc.user.Email, record.Email)
				mts.Require().Equal(tc.user.GithubUserID, record.GithubUserID)
				mts.Require().Equal(tc.user.GithubAccessToken, record.GithubAccessToken)
				mts.Require().Equal(tc.user.AvatarURL, record.AvatarURL)
				mts.Require().Equal(tc.user.GravatarID, record.GravatarID)
			}
		})
	}
}

func (mts *ModelsTestSuite) TestGetUserModules() {
	mts.resetDB()

	mod := models.Module{
		Name: "x/bank",
		Team: "cosmonauts",
		Owners: []models.User{
			{Name: "foo", Email: models.NewNullString("foo@cosmonauts.com")},
		},
		Authors: []models.User{
			{Name: "foo", Email: models.NewNullString("foo@cosmonauts.com")},
		},
		Version: models.ModuleVersion{
			Version:       "v1.0.0",
			Repo:          "https://github.com/cosmos/cosmos-sdk/releases/tag/v0.39.1",
			Documentation: "https://raw.githubusercontent.com/cosmos/cosmos-sdk/v0.39.1/x/bank/README.md",
		},
		Keywords: []models.Keyword{
			{Name: "tokens"}, {Name: "transfer"},
		},
		BugTracker: models.BugTracker{
			URL:     models.NewNullString("cosmonauts.com"),
			Contact: models.NewNullString("contact@cosmonauts.com"),
		},
	}

	record, err := mod.Upsert(mts.gormDB)
	mts.Require().NoError(err)

	mods, err := models.GetUserModules(mts.gormDB, record.Authors[0].Name)
	mts.Require().NoError(err)
	mts.Require().Len(mods, 1)
	mts.Require().Equal(mods[0], record)
}

func (mts *ModelsTestSuite) TestGetUserByID() {
	mts.resetDB()

	mod := models.Module{
		Name: "x/bank",
		Team: "cosmonauts",
		Authors: []models.User{
			{Name: "foo", Email: models.NewNullString("foo@cosmonauts.com")},
		},
		Version: models.ModuleVersion{
			Version:       "v1.0.0",
			Repo:          "https://github.com/cosmos/cosmos-sdk/releases/tag/v0.39.1",
			Documentation: "https://raw.githubusercontent.com/cosmos/cosmos-sdk/v0.39.1/x/bank/README.md",
		},
		Keywords: []models.Keyword{
			{Name: "tokens"}, {Name: "transfer"},
		},
		BugTracker: models.BugTracker{
			URL:     models.NewNullString("cosmonauts.com"),
			Contact: models.NewNullString("contact@cosmonauts.com"),
		},
	}

	mod, err := mod.Upsert(mts.gormDB)
	mts.Require().NoError(err)

	mts.Run("no user exists", func() {
		result, err := models.GetUserByID(mts.gormDB, mod.Authors[0].ID+1)
		mts.Require().Error(err)
		mts.Require().Equal(models.User{}, result)
	})

	mts.Run("user exists", func() {
		result, err := models.GetUserByID(mts.gormDB, mod.Authors[0].ID)
		mts.Require().NoError(err)
		mts.Require().Equal(mod.Authors[0].Name, result.Name)
		mts.Require().Equal(mod.Authors[0].Email, result.Email)
	})
}

func (mts *ModelsTestSuite) TestGetAllUsers() {
	mts.resetDB()

	users, paginator, err := models.GetAllUsers(mts.gormDB, httputil.PaginationQuery{Page: 1, Limit: 10, Order: "id"})
	mts.Require().NoError(err)
	mts.Require().Empty(users)
	mts.Require().Zero(paginator.PrevPage)
	mts.Require().Equal(int64(0), paginator.NextPage)

	for i := 0; i < 25; i++ {
		mod := models.Module{
			Name: fmt.Sprintf("x/bank-%d", i),
			Team: "cosmonauts",
			Authors: []models.User{
				{Name: fmt.Sprintf("foo-%d", i), Email: models.NewNullString(fmt.Sprintf("foo%d@cosmonauts.com", i))},
			},
			Version: models.ModuleVersion{
				Version:       "v1.0.0",
				Repo:          "https://github.com/cosmos/cosmos-sdk/releases/tag/v0.39.1",
				Documentation: fmt.Sprintf("https://raw.githubusercontent.com/cosmos/cosmos-sdk/v0.39.1/x/bank-%d/README.md", i),
			},
			Keywords: []models.Keyword{
				{Name: "tokens"},
			},
			BugTracker: models.BugTracker{
				URL:     models.NewNullString("cosmonauts.com"),
				Contact: models.NewNullString("contact@cosmonauts.com"),
			},
		}

		_, err := mod.Upsert(mts.gormDB)
		mts.Require().NoError(err)
	}

	// first page (full) ordered by newest
	users, paginator, err = models.GetAllUsers(mts.gormDB, httputil.PaginationQuery{Page: 1, Limit: 10, Order: "created_at,id", Reverse: true})
	mts.Require().NoError(err)
	mts.Require().Len(users, 10)
	mts.Require().Zero(paginator.PrevPage)
	mts.Require().Equal(int64(2), paginator.NextPage)
	mts.Require().Equal(uint(25), users[0].ID)
	mts.Require().Equal(uint(16), users[len(users)-1].ID)

	// first page (full)
	users, paginator, err = models.GetAllUsers(mts.gormDB, httputil.PaginationQuery{Page: 1, Limit: 10, Order: "id"})
	mts.Require().NoError(err)
	mts.Require().Len(users, 10)
	mts.Require().Zero(paginator.PrevPage)
	mts.Require().Equal(int64(2), paginator.NextPage)

	// second page (full)
	users, paginator, err = models.GetAllUsers(mts.gormDB, httputil.PaginationQuery{Page: 2, Limit: 10, Order: "id"})
	mts.Require().NoError(err)
	mts.Require().Len(users, 10)
	mts.Require().Equal(int64(1), paginator.PrevPage)
	mts.Require().Equal(int64(3), paginator.NextPage)

	// third page (partially full)
	users, paginator, err = models.GetAllUsers(mts.gormDB, httputil.PaginationQuery{Page: 3, Limit: 10, Order: "id"})
	mts.Require().NoError(err)
	mts.Require().Len(users, 5)
	mts.Require().Equal(int64(2), paginator.PrevPage)
	mts.Require().Zero(paginator.NextPage)
}

func (mts *ModelsTestSuite) TestGetAllKeywords() {
	mts.resetDB()

	keywords, paginator, err := models.GetAllKeywords(mts.gormDB, httputil.PaginationQuery{Page: 1, Limit: 10, Order: "id"})
	mts.Require().NoError(err)
	mts.Require().Empty(keywords)
	mts.Require().Zero(paginator.PrevPage)
	mts.Require().Equal(int64(0), paginator.NextPage)

	for i := 0; i < 25; i++ {
		mod := models.Module{
			Name: fmt.Sprintf("x/bank-%d", i),
			Team: "cosmonauts",
			Authors: []models.User{
				{Name: "foo", Email: models.NewNullString("foo@cosmonauts.com")},
			},
			Version: models.ModuleVersion{
				Version:       "v1.0.0",
				Repo:          "https://github.com/cosmos/cosmos-sdk/releases/tag/v0.39.1",
				Documentation: fmt.Sprintf("https://raw.githubusercontent.com/cosmos/cosmos-sdk/v0.39.1/x/bank-%d/README.md", i),
			},
			Keywords: []models.Keyword{
				{Name: fmt.Sprintf("tokens-%d", i)},
			},
			BugTracker: models.BugTracker{
				URL:     models.NewNullString("cosmonauts.com"),
				Contact: models.NewNullString("contact@cosmonauts.com"),
			},
		}

		_, err := mod.Upsert(mts.gormDB)
		mts.Require().NoError(err)
	}

	// first page (full) ordered by newest
	keywords, paginator, err = models.GetAllKeywords(mts.gormDB, httputil.PaginationQuery{Page: 1, Limit: 10, Order: "created_at,id", Reverse: true})
	mts.Require().NoError(err)
	mts.Require().Len(keywords, 10)
	mts.Require().Zero(paginator.PrevPage)
	mts.Require().Equal(int64(2), paginator.NextPage)
	mts.Require().Equal(uint(25), keywords[0].ID)
	mts.Require().Equal(uint(16), keywords[len(keywords)-1].ID)

	// first page (full)
	keywords, paginator, err = models.GetAllKeywords(mts.gormDB, httputil.PaginationQuery{Page: 1, Limit: 10, Order: "id"})
	mts.Require().NoError(err)
	mts.Require().Len(keywords, 10)
	mts.Require().Zero(paginator.PrevPage)
	mts.Require().Equal(int64(2), paginator.NextPage)

	// second page (full)
	keywords, paginator, err = models.GetAllKeywords(mts.gormDB, httputil.PaginationQuery{Page: 2, Limit: 10, Order: "id"})
	mts.Require().NoError(err)
	mts.Require().Len(keywords, 10)
	mts.Require().Equal(int64(1), paginator.PrevPage)
	mts.Require().Equal(int64(3), paginator.NextPage)

	// third page (partially full)
	keywords, paginator, err = models.GetAllKeywords(mts.gormDB, httputil.PaginationQuery{Page: 3, Limit: 10, Order: "id"})
	mts.Require().NoError(err)
	mts.Require().Len(keywords, 5)
	mts.Require().Equal(int64(2), paginator.PrevPage)
	mts.Require().Zero(paginator.NextPage)
}

func (mts *ModelsTestSuite) TestModuleStar() {
	mts.resetDB()

	mod := models.Module{
		Name: "x/bank",
		Team: "cosmonauts",
		Owners: []models.User{
			{Name: "foo", Email: models.NewNullString("foo@cosmonauts.com")},
		},
		Authors: []models.User{
			{Name: "foo", Email: models.NewNullString("foo@cosmonauts.com")},
		},
		Version: models.ModuleVersion{
			Version:       "v1.0.0",
			Repo:          "https://github.com/cosmos/cosmos-sdk/releases/tag/v0.39.1",
			Documentation: "https://raw.githubusercontent.com/cosmos/cosmos-sdk/v0.39.1/x/bank/README.md",
		},
		Keywords: []models.Keyword{
			{Name: "tokens"}, {Name: "transfer"},
		},
		BugTracker: models.BugTracker{
			URL:     models.NewNullString("cosmonauts.com"),
			Contact: models.NewNullString("contact@cosmonauts.com"),
		},
	}

	// create module
	mod, err := mod.Upsert(mts.gormDB)
	mts.Require().NoError(err)

	// ensure the owner has not favored it
	ok, err := mod.UserStarred(mts.gormDB, mod.Owners[0].ID)
	mts.Require().NoError(err)
	mts.Require().False(ok)

	// ensure we can favorite it and the count matches
	stars, err := mod.Star(mts.gormDB, mod.Owners[0].ID)
	mts.Require().NoError(err)
	mts.Require().Equal(int64(1), stars)

	// ensure the owner has favored the module
	ok, err = mod.UserStarred(mts.gormDB, mod.Owners[0].ID)
	mts.Require().NoError(err)
	mts.Require().True(ok)

	// ensure we can un-favorite it and the count matches
	stars, err = mod.UnStar(mts.gormDB, mod.Owners[0].ID)
	mts.Require().NoError(err)
	mts.Require().Equal(int64(0), stars)

	// ensure the owner has not favored it
	ok, err = mod.UserStarred(mts.gormDB, mod.Owners[0].ID)
	mts.Require().NoError(err)
	mts.Require().False(ok)
}

func (mts *ModelsTestSuite) TestUserEmailConfirmation_Upsert() {
	mts.resetDB()

	mod := models.Module{
		Name: "x/bank",
		Team: "cosmonauts",
		Owners: []models.User{
			{Name: "foo"},
		},
		Authors: []models.User{
			{Name: "foo"},
		},
		Version: models.ModuleVersion{
			Version:       "v1.0.0",
			Repo:          "https://github.com/cosmos/cosmos-sdk/releases/tag/v0.39.1",
			Documentation: "https://raw.githubusercontent.com/cosmos/cosmos-sdk/v0.39.1/x/bank/README.md",
		},
		Keywords: []models.Keyword{
			{Name: "tokens"}, {Name: "transfer"},
		},
		BugTracker: models.BugTracker{
			URL:     models.NewNullString("cosmonauts.com"),
			Contact: models.NewNullString("contact@cosmonauts.com"),
		},
	}

	// create module
	mod, err := mod.Upsert(mts.gormDB)
	mts.Require().NoError(err)

	user, err := models.GetUserByID(mts.gormDB, mod.Owners[0].ID)
	mts.Require().NoError(err)
	mts.Require().False(user.EmailConfirmed)

	uec, err := models.UserEmailConfirmation{UserID: mod.Owners[0].ID, Email: "foo@email.com"}.Upsert(mts.gormDB)
	token1 := uec.Token
	mts.Require().NoError(err)
	mts.Require().NotEqual(uuid.UUID{}, token1)
	mts.Require().Equal("foo@email.com", uec.Email)

	uec, err = models.UserEmailConfirmation{UserID: mod.Owners[0].ID, Email: "bar@email.com"}.Upsert(mts.gormDB)
	token2 := uec.Token
	mts.Require().NoError(err)
	mts.Require().NotEqual(uuid.UUID{}, token2)
	mts.Require().NotEqual(token1, token2)
	mts.Require().Equal("bar@email.com", uec.Email)

	var count int64
	mts.Require().NoError(mts.gormDB.Table("user_email_confirmations").Count(&count).Error)
	mts.Require().Equal(int64(1), count)
}

func (mts *ModelsTestSuite) TestQueryUserEmailConfirmation() {
	mts.resetDB()

	mod := models.Module{
		Name: "x/bank",
		Team: "cosmonauts",
		Owners: []models.User{
			{Name: "foo", Email: models.NewNullString("foo@cosmonauts.com")},
		},
		Authors: []models.User{
			{Name: "foo", Email: models.NewNullString("foo@cosmonauts.com")},
		},
		Version: models.ModuleVersion{
			Version:       "v1.0.0",
			Repo:          "https://github.com/cosmos/cosmos-sdk/releases/tag/v0.39.1",
			Documentation: "https://raw.githubusercontent.com/cosmos/cosmos-sdk/v0.39.1/x/bank/README.md",
		},
		Keywords: []models.Keyword{
			{Name: "tokens"}, {Name: "transfer"},
		},
		BugTracker: models.BugTracker{
			URL:     models.NewNullString("cosmonauts.com"),
			Contact: models.NewNullString("contact@cosmonauts.com"),
		},
	}

	// create module
	mod, err := mod.Upsert(mts.gormDB)
	mts.Require().NoError(err)

	user, err := models.GetUserByID(mts.gormDB, mod.Owners[0].ID)
	mts.Require().NoError(err)
	mts.Require().False(user.EmailConfirmed)

	query := map[string]interface{}{"user_id": mod.Owners[0].ID}

	uec, err := models.QueryUserEmailConfirmation(mts.gormDB, query)
	mts.Require().Error(err)
	mts.Require().Empty(uec)

	uec, err = models.UserEmailConfirmation{UserID: mod.Owners[0].ID, Email: "foo@email.com"}.Upsert(mts.gormDB)
	token := uec.Token
	mts.Require().NoError(err)
	mts.Require().NotEqual(uuid.UUID{}, token)
	mts.Require().Equal("foo@email.com", uec.Email)

	uec, err = models.QueryUserEmailConfirmation(mts.gormDB, query)
	mts.Require().NoError(err)
	mts.Require().Equal(mod.Owners[0].ID, uec.UserID)
	mts.Require().Equal(token, uec.Token)
	mts.Require().Equal("foo@email.com", uec.Email)
}

func (mts *ModelsTestSuite) TestUser_ConfirmEmail() {
	mts.resetDB()

	mod := models.Module{
		Name: "x/bank",
		Team: "cosmonauts",
		Owners: []models.User{
			{Name: "foo", Email: models.NewNullString("foo@cosmonauts.com")},
		},
		Authors: []models.User{
			{Name: "foo", Email: models.NewNullString("foo@cosmonauts.com")},
		},
		Version: models.ModuleVersion{
			Version:       "v1.0.0",
			Repo:          "https://github.com/cosmos/cosmos-sdk/releases/tag/v0.39.1",
			Documentation: "https://raw.githubusercontent.com/cosmos/cosmos-sdk/v0.39.1/x/bank/README.md",
		},
		Keywords: []models.Keyword{
			{Name: "tokens"}, {Name: "transfer"},
		},
		BugTracker: models.BugTracker{
			URL:     models.NewNullString("cosmonauts.com"),
			Contact: models.NewNullString("contact@cosmonauts.com"),
		},
	}

	// create module
	mod, err := mod.Upsert(mts.gormDB)
	mts.Require().NoError(err)

	user, err := models.GetUserByID(mts.gormDB, mod.Owners[0].ID)
	mts.Require().NoError(err)
	mts.Require().False(user.EmailConfirmed)
	mts.Require().Equal("foo@cosmonauts.com", user.Email.String)

	uec, err := models.UserEmailConfirmation{UserID: mod.Owners[0].ID, Email: "foo@email.com"}.Upsert(mts.gormDB)
	token := uec.Token
	mts.Require().NoError(err)
	mts.Require().NotEqual(uuid.UUID{}, token)
	mts.Require().Equal("foo@email.com", uec.Email)

	user, err = user.ConfirmEmail(mts.gormDB, uec)
	mts.Require().NoError(err)
	mts.Require().True(user.EmailConfirmed)
	mts.Require().Equal("foo@email.com", user.Email.String)

	query := map[string]interface{}{"user_id": mod.Owners[0].ID}
	uec, err = models.QueryUserEmailConfirmation(mts.gormDB, query)
	mts.Require().Error(err)
	mts.Require().Empty(uec)
}

func (mts *ModelsTestSuite) TestModuleOwnerInvite_Upsert() {
	mts.resetDB()

	mod := models.Module{
		Name: "x/bank",
		Team: "cosmonauts",
		Owners: []models.User{
			{Name: "foo"},
		},
		Authors: []models.User{
			{Name: "bar"},
		},
		Version: models.ModuleVersion{
			Version:       "v1.0.0",
			Repo:          "https://github.com/cosmos/cosmos-sdk/releases/tag/v0.39.1",
			Documentation: "https://raw.githubusercontent.com/cosmos/cosmos-sdk/v0.39.1/x/bank/README.md",
		},
		Keywords: []models.Keyword{
			{Name: "tokens"}, {Name: "transfer"},
		},
		BugTracker: models.BugTracker{
			URL:     models.NewNullString("cosmonauts.com"),
			Contact: models.NewNullString("contact@cosmonauts.com"),
		},
	}

	// create module
	mod, err := mod.Upsert(mts.gormDB)
	mts.Require().NoError(err)

	owner, err := models.GetUserByID(mts.gormDB, mod.Owners[0].ID)
	mts.Require().NoError(err)

	author, err := models.GetUserByID(mts.gormDB, mod.Authors[0].ID)
	mts.Require().NoError(err)

	moi, err := models.ModuleOwnerInvite{ModuleID: mod.ID, InvitedByUserID: owner.ID, InvitedUserID: author.ID}.Upsert(mts.gormDB)
	token1 := moi.Token
	mts.Require().NoError(err)
	mts.Require().NotEqual(uuid.UUID{}, token1)

	moi, err = models.ModuleOwnerInvite{ModuleID: mod.ID, InvitedByUserID: owner.ID, InvitedUserID: author.ID}.Upsert(mts.gormDB)
	token2 := moi.Token
	mts.Require().NoError(err)
	mts.Require().NotEqual(uuid.UUID{}, token2)
	mts.Require().NotEqual(token1, token2)

	var count int64
	mts.Require().NoError(mts.gormDB.Table("module_owner_invites").Count(&count).Error)
	mts.Require().Equal(int64(1), count)
}

func (mts *ModelsTestSuite) TestQueryModuleOwnerInvite() {
	mts.resetDB()

	mod := models.Module{
		Name: "x/bank",
		Team: "cosmonauts",
		Owners: []models.User{
			{Name: "foo"},
		},
		Authors: []models.User{
			{Name: "bar"},
		},
		Version: models.ModuleVersion{
			Version:       "v1.0.0",
			Repo:          "https://github.com/cosmos/cosmos-sdk/releases/tag/v0.39.1",
			Documentation: "https://raw.githubusercontent.com/cosmos/cosmos-sdk/v0.39.1/x/bank/README.md",
		},
		Keywords: []models.Keyword{
			{Name: "tokens"}, {Name: "transfer"},
		},
		BugTracker: models.BugTracker{
			URL:     models.NewNullString("cosmonauts.com"),
			Contact: models.NewNullString("contact@cosmonauts.com"),
		},
	}

	// create module
	mod, err := mod.Upsert(mts.gormDB)
	mts.Require().NoError(err)

	owner, err := models.GetUserByID(mts.gormDB, mod.Owners[0].ID)
	mts.Require().NoError(err)

	author, err := models.GetUserByID(mts.gormDB, mod.Authors[0].ID)
	mts.Require().NoError(err)

	query := map[string]interface{}{"module_id": mod.ID, "invited_user_id": author.ID, "invited_by_user_id": owner.ID}

	moi, err := models.QueryModuleOwnerInvite(mts.gormDB, query)
	mts.Require().Error(err)
	mts.Require().Empty(moi)

	moi, err = models.ModuleOwnerInvite{ModuleID: mod.ID, InvitedByUserID: owner.ID, InvitedUserID: author.ID}.Upsert(mts.gormDB)
	token := moi.Token
	mts.Require().NoError(err)
	mts.Require().NotEqual(uuid.UUID{}, token)

	moi, err = models.QueryModuleOwnerInvite(mts.gormDB, query)
	mts.Require().NoError(err)
	mts.Require().Equal(mod.ID, moi.ModuleID)
	mts.Require().Equal(author.ID, moi.InvitedUserID)
	mts.Require().Equal(owner.ID, moi.InvitedByUserID)
	mts.Require().Equal(token, moi.Token)
}

func (mts *ModelsTestSuite) TestModule_AddOwner() {
	mts.resetDB()

	mod := models.Module{
		Name: "x/bank",
		Team: "cosmonauts",
		Owners: []models.User{
			{Name: "foo"},
		},
		Authors: []models.User{
			{Name: "bar"},
		},
		Version: models.ModuleVersion{
			Version:       "v1.0.0",
			Repo:          "https://github.com/cosmos/cosmos-sdk/releases/tag/v0.39.1",
			Documentation: "https://raw.githubusercontent.com/cosmos/cosmos-sdk/v0.39.1/x/bank/README.md",
		},
		Keywords: []models.Keyword{
			{Name: "tokens"}, {Name: "transfer"},
		},
		BugTracker: models.BugTracker{
			URL:     models.NewNullString("cosmonauts.com"),
			Contact: models.NewNullString("contact@cosmonauts.com"),
		},
	}

	// create module
	mod, err := mod.Upsert(mts.gormDB)
	mts.Require().NoError(err)

	owner, err := models.GetUserByID(mts.gormDB, mod.Owners[0].ID)
	mts.Require().NoError(err)

	author, err := models.GetUserByID(mts.gormDB, mod.Authors[0].ID)
	mts.Require().NoError(err)

	moi, err := models.ModuleOwnerInvite{ModuleID: mod.ID, InvitedByUserID: owner.ID, InvitedUserID: author.ID}.Upsert(mts.gormDB)
	token := moi.Token
	mts.Require().NoError(err)
	mts.Require().NotEqual(uuid.UUID{}, token)

	mod, err = mod.AddOwner(mts.gormDB, author)
	mts.Require().NoError(err)
	mts.Require().Len(mod.Owners, 2)

	ownerNames := make([]string, len(mod.Owners))
	for i, o := range mod.Owners {
		ownerNames[i] = o.Name
	}
	mts.Require().Contains(ownerNames, author.Name)

	query := map[string]interface{}{"module_id": mod.ID, "invited_user_id": author.ID, "invited_by_user_id": owner.ID}
	moi, err = models.QueryModuleOwnerInvite(mts.gormDB, query)
	mts.Require().Error(err)
	mts.Require().Empty(moi)
}

func (mts *ModelsTestSuite) TestNewBugTrackerJSON() {
	bugTracker := models.BugTracker{
		Model: gorm.Model{
			ID:        1,
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
		},
		URL:     models.NewNullString("url"),
		Contact: models.NewNullString("contact"),
	}
	bugTrackerJSON := bugTracker.NewBugTrackerJSON()
	mts.Require().Equal(bugTracker.URL.String, bugTrackerJSON.URL)
	mts.Require().Equal(bugTracker.Contact.String, bugTrackerJSON.Contact)
	mts.Require().Equal(bugTracker.CreatedAt, bugTrackerJSON.CreatedAt)
	mts.Require().Equal(bugTracker.UpdatedAt, bugTrackerJSON.UpdatedAt)
}

func (mts *ModelsTestSuite) TestNewModuleVersionJSON() {
	version := models.ModuleVersion{
		Model: gorm.Model{
			ID:        1,
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
		},
		Documentation: "documentation",
		Repo:          "repo",
		Version:       "version",
		SDKCompat:     models.NewNullString("sdk compat"),
		ModuleID:      1,
		PublishedBy:   1,
	}
	versionJSON := version.NewModuleVersionJSON()
	mts.Require().Equal(version.Documentation, versionJSON.Documentation)
	mts.Require().Equal(version.Repo, versionJSON.Repo)
	mts.Require().Equal(version.Version, versionJSON.Version)
	mts.Require().Equal(version.SDKCompat.String, versionJSON.SDKCompat)
	mts.Require().Equal(version.ModuleID, versionJSON.ModuleID)
	mts.Require().Equal(version.PublishedBy, versionJSON.PublishedBy)
	mts.Require().Equal(version.CreatedAt, versionJSON.CreatedAt)
	mts.Require().Equal(version.UpdatedAt, versionJSON.UpdatedAt)
}

func (mts *ModelsTestSuite) TestNewKeywordJSON() {
	keyword := models.Keyword{
		Model: gorm.Model{
			ID:        1,
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
		},
		Name: "keyword",
	}
	keywordJSON := keyword.NewKeywordJSON()
	mts.Require().Equal(keyword.Name, keywordJSON.Name)
	mts.Require().Equal(keyword.CreatedAt, keywordJSON.CreatedAt)
	mts.Require().Equal(keyword.UpdatedAt, keywordJSON.UpdatedAt)
}

func (mts *ModelsTestSuite) TestNewUserJSON() {
	user := models.User{
		Model: gorm.Model{
			ID:        1,
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
		},
		Name:           "user",
		FullName:       "full name",
		URL:            "url",
		GravatarID:     "gravatar ID",
		AvatarURL:      "avatar ID",
		EmailConfirmed: true,
	}
	userJSON := user.NewUserJSON()
	mts.Require().Equal(user.Name, userJSON.Name)
	mts.Require().Equal(user.CreatedAt, userJSON.CreatedAt)
	mts.Require().Equal(user.UpdatedAt, userJSON.UpdatedAt)
	mts.Require().Equal(user.FullName, userJSON.FullName)
	mts.Require().Equal(user.URL, userJSON.URL)
	mts.Require().Equal(user.GravatarID, userJSON.GravatarID)
	mts.Require().Equal(user.AvatarURL, userJSON.AvatarURL)
	mts.Require().Equal(user.EmailConfirmed, userJSON.EmailConfirmed)
}

func (mts *ModelsTestSuite) TestLocationUpsert() {
	mts.resetDB()

	testCases := []struct {
		name      string
		loc       models.Location
		expectErr bool
	}{
		{
			"valid location",
			models.Location{
				Country:   "US",
				Region:    "US",
				City:      "New York",
				Latitude:  "40.7128",
				Longitude: "74.0060",
			},
			false,
		},
		{
			"empty region",
			models.Location{
				Country:   "US",
				City:      "New York",
				Latitude:  "40.7128",
				Longitude: "74.0060",
			},
			false,
		},
		{
			"updated region",
			models.Location{
				Country:   "US",
				City:      "New York",
				Region:    "new region",
				Latitude:  "40.7128",
				Longitude: "74.0060",
			},
			false,
		},
		{
			"missing latitude",
			models.Location{
				Country:   "US",
				City:      "New York",
				Region:    "new region",
				Longitude: "74.0060",
			},
			true,
		},
		{
			"missing longitude",
			models.Location{
				Country:  "US",
				City:     "New York",
				Region:   "new region",
				Latitude: "40.7128",
			},
			true,
		},
	}

	for _, tc := range testCases {
		tc := tc

		mts.Run(tc.name, func() {
			record, err := tc.loc.Upsert(mts.gormDB)
			if tc.expectErr {
				mts.Require().Error(err)
			} else {
				mts.Require().NoError(err)
				mts.Require().Equal(tc.loc.Country, record.Country)
				mts.Require().Equal(tc.loc.Region, record.Region)
				mts.Require().Equal(tc.loc.City, record.City)
				mts.Require().Equal(tc.loc.Latitude, record.Latitude)
				mts.Require().Equal(tc.loc.Longitude, record.Longitude)
			}
		})
	}
}

func (mts *ModelsTestSuite) TestNewLocationJSON() {
	loc := models.Location{
		ID:        1,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Country:   "US",
		Region:    "US",
		City:      "New York",
		Latitude:  "40.7128",
		Longitude: "74.0060",
	}

	locJSON := loc.NewLocationJSON()
	mts.Require().Equal(loc.Country, locJSON.Country)
	mts.Require().Equal(loc.Region, locJSON.Region)
	mts.Require().Equal(loc.City, locJSON.City)
	mts.Require().Equal(loc.Latitude, locJSON.Latitude)
	mts.Require().Equal(loc.Longitude, locJSON.Longitude)
	mts.Require().Equal(loc.ID, loc.ID)
	mts.Require().Equal(loc.CreatedAt, loc.CreatedAt)
	mts.Require().Equal(loc.UpdatedAt, loc.UpdatedAt)
}

func (mts *ModelsTestSuite) TestNodeUpsert() {
	mts.resetDB()

	testCases := []struct {
		name      string
		node      models.Node
		expectErr bool
	}{
		{
			"valid node",
			models.Node{
				Location: models.Location{
					Country:   "US",
					Region:    "US",
					City:      "New York",
					Latitude:  "40.7128",
					Longitude: "74.0060",
				},
				Address: "127.0.0.1",
				RPCPort: "26657",
				P2PPort: "26656",
				Moniker: "test",
				NodeID:  "0000FF",
				Network: "testnet",
				Version: "1.0.1",
				TxIndex: "false",
			},
			false,
		},
		{
			"updated location",
			models.Node{
				Location: models.Location{
					Country:   "US",
					Region:    "US",
					City:      "Baltimore",
					Latitude:  "33.7128",
					Longitude: "25.0060",
				},
				Address: "127.0.0.1",
				RPCPort: "26657",
				P2PPort: "26656",
				Moniker: "test",
				NodeID:  "0000FF",
				Network: "testnet",
				Version: "1.0.1",
				TxIndex: "false",
			},
			false,
		},
		{
			"same address for different network",
			models.Node{
				Location: models.Location{
					Country:   "US",
					Region:    "US",
					City:      "New York",
					Latitude:  "40.7128",
					Longitude: "74.0060",
				},
				Address: "127.0.0.1",
				RPCPort: "26657",
				P2PPort: "26656",
				Moniker: "test",
				NodeID:  "0000FF",
				Network: "other",
				Version: "1.0.1",
				TxIndex: "false",
			},
			false,
		},
	}

	for _, tc := range testCases {
		tc := tc

		mts.Run(tc.name, func() {
			record, err := tc.node.Upsert(mts.gormDB)
			if tc.expectErr {
				mts.Require().Error(err)
			} else {
				mts.Require().NoError(err)
				mts.Require().Equal(tc.node.Address, record.Address)
				mts.Require().Equal(tc.node.RPCPort, record.RPCPort)
				mts.Require().Equal(tc.node.P2PPort, record.P2PPort)
				mts.Require().Equal(tc.node.Moniker, record.Moniker)
				mts.Require().Equal(tc.node.NodeID, record.NodeID)
				mts.Require().Equal(tc.node.Network, record.Network)
				mts.Require().Equal(tc.node.Version, record.Version)
				mts.Require().Equal(tc.node.TxIndex, record.TxIndex)
				mts.Require().Equal(tc.node.Location.Country, record.Location.Country)
				mts.Require().Equal(tc.node.Location.Region, record.Location.Region)
				mts.Require().Equal(tc.node.Location.City, record.Location.City)
				mts.Require().Equal(tc.node.Location.Latitude, record.Location.Latitude)
				mts.Require().Equal(tc.node.Location.Longitude, record.Location.Longitude)
			}
		})
	}
}

func (mts *ModelsTestSuite) TestNewNodeJSON() {
	node := models.Node{
		ID:        1,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Address:   "127.0.0.1",
		RPCPort:   "25626",
		P2PPort:   "25627",
		Moniker:   "foo",
		NodeID:    "00000FF",
		Network:   "testnet-3",
		Version:   "1.0",
		TxIndex:   "false",
	}

	nodeJSON := node.NewNodeJSON()
	mts.Require().Equal(nodeJSON.ID, nodeJSON.ID)
	mts.Require().Equal(nodeJSON.CreatedAt, nodeJSON.CreatedAt)
	mts.Require().Equal(nodeJSON.UpdatedAt, nodeJSON.UpdatedAt)
	mts.Require().Equal(node.Address, nodeJSON.Address)
	mts.Require().Equal(node.RPCPort, nodeJSON.RPCPort)
	mts.Require().Equal(node.P2PPort, nodeJSON.P2PPort)
	mts.Require().Equal(node.Moniker, nodeJSON.Moniker)
	mts.Require().Equal(node.NodeID, nodeJSON.NodeID)
	mts.Require().Equal(node.Network, nodeJSON.Network)
	mts.Require().Equal(node.Version, nodeJSON.Version)
	mts.Require().Equal(node.TxIndex, nodeJSON.TxIndex)
}

func (mts *ModelsTestSuite) TestNodeDelete() {
	mts.resetDB()

	node := models.Node{}
	mts.Require().NoError(node.Delete(mts.gormDB))

	node = models.Node{
		Location: models.Location{
			Country:   "US",
			Region:    "US",
			City:      "New York",
			Latitude:  "40.7128",
			Longitude: "74.0060",
		},
		Address: "127.0.0.1",
		RPCPort: "26657",
		P2PPort: "26656",
		Moniker: "test",
		NodeID:  "0000FF",
		Network: "testnet",
		Version: "1.0.1",
		TxIndex: "false",
	}

	_, err := node.Upsert(mts.gormDB)
	mts.Require().NoError(err)
	mts.Require().NoError(node.Delete(mts.gormDB))

	_, err = models.QueryNode(
		mts.gormDB,
		map[string]interface{}{"address": node.Address, "network": node.Network},
	)
	mts.Require().Error(err)
}

func (mts *ModelsTestSuite) TestGetStaleNodes() {
	mts.resetDB()

	var last time.Time

	for i := 0; i < 10; i++ {
		node := models.Node{
			Location: models.Location{
				Country:   "US",
				Region:    "US",
				City:      "New York",
				Latitude:  "40.7128",
				Longitude: "74.0060",
			},
			Address: fmt.Sprintf("127.0.0.%d", i),
			RPCPort: "26657",
			P2PPort: "26656",
			Moniker: fmt.Sprintf("test-node-%d", i),
			NodeID:  "0000FF",
			Network: "testnet",
			Version: "1.0.1",
			TxIndex: "false",
		}

		record, err := node.Upsert(mts.gormDB)
		mts.Require().NoError(err)

		last = record.UpdatedAt
	}

	nodes, err := models.GetStaleNodes(mts.gormDB, last.Add(time.Minute))
	mts.Require().NoError(err)
	mts.Require().Len(nodes, 10)

	nodes, err = models.GetStaleNodes(mts.gormDB, last.Add(-10*time.Minute))
	mts.Require().NoError(err)
	mts.Require().Empty(nodes)
}

func (mts *ModelsTestSuite) resetDB() {
	mts.T().Helper()

	if err := mts.m.Down(); err != nil {
		require.Equal(mts.T(), migrate.ErrNoChange, err, fmt.Sprintf("down migration error: %s", err.Error()))
	}
	if err := mts.m.Up(); err != nil {
		require.Equal(mts.T(), migrate.ErrNoChange, err, fmt.Sprintf("up migration error: %s", err.Error()))
	}
}
