package models_test

import (
	"database/sql"
	"fmt"
	"math/rand"
	"os"
	"testing"

	"github.com/golang-migrate/migrate/v4"
	migratepg "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"

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

	gormDB, err := gorm.Open(postgres.Open(connStr), &gorm.Config{Logger: gormlogger.Discard})
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
	resetDB(mts.T(), mts.m)

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
				Repo: "https://github.com/cosmos/cosmos-sdk",
			},
			expectErr: true,
		},
		{
			name: "create module no version",
			module: models.Module{
				Name: "x/bank",
				Team: "cosmonauts",
				Repo: "https://github.com/cosmos/cosmos-sdk",
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
				Repo: "https://github.com/cosmos/cosmos-sdk",
				Authors: []models.User{
					{Name: "foo", Email: models.NewNullString("foo@cosmonauts.com")},
				},
				Version: "v1.0.0",
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
			result, err := tc.module.Upsert(mts.gormDB)
			if tc.expectErr {
				mts.Require().Error(err)
			} else {
				mts.Require().NoError(err)
				mts.Require().Equal(tc.module.Name, result.Name)
				mts.Require().Equal(tc.module.Team, result.Team)
				mts.Require().Equal(tc.module.Description, result.Description)
				mts.Require().Equal(tc.module.Homepage, result.Homepage)
				mts.Require().Equal(tc.module.Documentation, result.Documentation)
			}
		})
	}
}

func (mts *ModelsTestSuite) TestGetModuleByID() {
	resetDB(mts.T(), mts.m)

	mod := models.Module{
		Name: "x/bank",
		Team: "cosmonauts",
		Repo: "https://github.com/cosmos/cosmos-sdk",
		Authors: []models.User{
			{Name: "foo", Email: models.NewNullString("foo@cosmonauts.com")},
		},
		Version: "v1.0.0",
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
		mts.Require().Equal(mod.Documentation, result.Documentation)
	})
}

func (mts *ModelsTestSuite) TestGetAllModules() {
	resetDB(mts.T(), mts.m)

	mods, err := models.GetAllModules(mts.gormDB, 0, 10)
	mts.Require().NoError(err)
	mts.Require().Empty(mods)

	for i := 0; i < 25; i++ {
		mod := models.Module{
			Name: fmt.Sprintf("x/bank-%d", i),
			Team: "cosmonauts",
			Repo: "https://github.com/cosmos/cosmos-sdk",
			Authors: []models.User{
				{Name: "foo", Email: models.NewNullString("foo@cosmonauts.com")},
			},
			Version: "v1.0.0",
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

	// first page (full)
	mods, err = models.GetAllModules(mts.gormDB, 0, 10)
	mts.Require().NoError(err)
	mts.Require().Len(mods, 10)

	cursor := mods[len(mods)-1].ID
	mts.Require().Equal(uint(10), cursor)

	// second page (full)
	mods, err = models.GetAllModules(mts.gormDB, cursor, 10)
	mts.Require().NoError(err)
	mts.Require().Len(mods, 10)

	cursor = mods[len(mods)-1].ID
	mts.Require().Equal(uint(20), cursor)

	// third page (partially full)
	mods, err = models.GetAllModules(mts.gormDB, cursor, 10)
	mts.Require().NoError(err)
	mts.Require().Len(mods, 5)

	cursor = mods[len(mods)-1].ID
	mts.Require().Equal(uint(25), cursor)
}

func (mts *ModelsTestSuite) TestModuleUpdateBasic() {
	resetDB(mts.T(), mts.m)

	mod := models.Module{
		Name:          "x/bank",
		Team:          "cosmonauts",
		Description:   "test description",
		Documentation: "https://github.com/cosmos/cosmos-sdk/x/bank/old_readme.md",
		Homepage:      "https://old.cosmos.network",
		Repo:          "https://github.com/cosmos/cosmos-sdk/old",
		Authors: []models.User{
			{Name: "admin"},
		},
		Version:    "v1.0.0",
		BugTracker: models.BugTracker{},
	}

	record, err := mod.Upsert(mts.gormDB)
	mts.Require().NoError(err)

	mod = models.Module{
		Name:          record.Name,
		Team:          record.Team,
		Description:   "new test description",
		Documentation: "https://github.com/cosmos/cosmos-sdk/x/bank/new_readme.md",
		Homepage:      "https://new.cosmos.network",
		Repo:          "https://github.com/cosmos/cosmos-sdk/new",
	}

	record, err = mod.Upsert(mts.gormDB)
	mts.Require().NoError(err)
	mts.Require().Equal(mod.Description, record.Description)
	mts.Require().Equal(mod.Documentation, record.Documentation)
	mts.Require().Equal(mod.Homepage, record.Homepage)
	mts.Require().Equal(mod.Repo, record.Repo)
}

func (mts *ModelsTestSuite) TestModuleUpdateBugTracker() {
	resetDB(mts.T(), mts.m)

	mod := models.Module{
		Name: "x/bank",
		Team: "cosmonauts",
		Repo: "https://github.com/cosmos/cosmos-sdk",
		Authors: []models.User{
			{Name: "admin"},
		},
		Version:    "v1.0.0",
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
	resetDB(mts.T(), mts.m)

	mod := models.Module{
		Name: "x/bank",
		Team: "cosmonauts",
		Repo: "https://github.com/cosmos/cosmos-sdk",
		Authors: []models.User{
			{Name: "admin"},
		},
		Version:    "v1.0.0",
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
	resetDB(mts.T(), mts.m)

	mod := models.Module{
		Name: "x/bank",
		Team: "cosmonauts",
		Repo: "https://github.com/cosmos/cosmos-sdk",
		Authors: []models.User{
			{Name: "admin"},
		},
		Owners: []models.User{
			{Name: "admin"},
		},
		Version:    "v1.0.0",
		BugTracker: models.BugTracker{},
	}

	record, err := mod.Upsert(mts.gormDB)
	mts.Require().NoError(err)

	mod.Authors = []models.User{
		{Name: "admin"}, {Name: "user1"}, {Name: "user2"},
	}

	record, err = mod.Upsert(mts.gormDB)
	mts.Require().NoError(err)

	for i, a := range record.Authors {
		fmt.Println("record created at:", a.CreatedAt)
		fmt.Println("module created at:", mod.Authors[i].CreatedAt)
	}

	mts.Require().Equal(mod.Authors, record.Authors)
	mts.Require().Equal(mod.Owners, record.Owners)

	mod.Authors = []models.User{
		{Name: "admin"}, {Name: "user2"},
	}

	record, err = mod.Upsert(mts.gormDB)
	mts.Require().NoError(err)
	mts.Require().Equal(mod.Authors, record.Authors)
	mts.Require().Equal(mod.Owners, record.Owners)

	mod.Authors = []models.User{}

	record, err = mod.Upsert(mts.gormDB)
	mts.Require().NoError(err)
	mts.Require().Equal(mod.Authors, record.Authors)
	mts.Require().Equal(mod.Owners, record.Owners)
}

func (mts *ModelsTestSuite) TestModuleUpdateOwners() {
	resetDB(mts.T(), mts.m)

	mod := models.Module{
		Name: "x/bank",
		Team: "cosmonauts",
		Repo: "https://github.com/cosmos/cosmos-sdk",
		Authors: []models.User{
			{Name: "admin"},
		},
		Owners: []models.User{
			{Name: "admin"},
		},
		Version:    "v1.0.0",
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
	mts.Require().Equal(mod.Owners, record.Owners)

	mod.Owners = []models.User{
		{Name: "admin"}, {Name: "user2"},
	}

	record, err = mod.Upsert(mts.gormDB)
	mts.Require().NoError(err)
	mts.Require().Equal(mod.Authors, record.Authors)
	mts.Require().Equal(mod.Owners, record.Owners)

	mod.Owners = []models.User{}

	record, err = mod.Upsert(mts.gormDB)
	mts.Require().NoError(err)
	mts.Require().Equal(mod.Authors, record.Authors)
	mts.Require().Equal(mod.Owners, record.Owners)
}

func (mts *ModelsTestSuite) TestModuleUpdateVersion() {
	resetDB(mts.T(), mts.m)

	mod := models.Module{
		Name: "x/bank",
		Team: "cosmonauts",
		Repo: "https://github.com/cosmos/cosmos-sdk",
		Authors: []models.User{
			{Name: "admin"},
		},
		Owners: []models.User{
			{Name: "admin"},
		},
		Version:    "v1.0.0",
		BugTracker: models.BugTracker{},
	}

	record, err := mod.Upsert(mts.gormDB)
	mts.Require().NoError(err)

	// update version
	mod.Version = "v1.0.1"

	record, err = mod.Upsert(mts.gormDB)
	mts.Require().NoError(err)
	mts.Require().Len(record.Versions, 2)

	latest, err := record.GetLatestVersion(mts.gormDB)
	mts.Require().NoError(err)
	mts.Require().Equal(mod.Version, latest.Version)

	// no version update
	mod.Version = "v1.0.1"

	record, err = mod.Upsert(mts.gormDB)
	mts.Require().NoError(err)
	mts.Require().Len(record.Versions, 2)

	latest, err = record.GetLatestVersion(mts.gormDB)
	mts.Require().NoError(err)
	mts.Require().Equal(mod.Version, latest.Version)

	// update version again
	mod.Version = "v2.0.0"

	record, err = mod.Upsert(mts.gormDB)
	mts.Require().NoError(err)
	mts.Require().Len(record.Versions, 3)

	latest, err = record.GetLatestVersion(mts.gormDB)
	mts.Require().NoError(err)
	mts.Require().Equal(mod.Version, latest.Version)
}

func (mts *ModelsTestSuite) TestModuleSearch() {
	resetDB(mts.T(), mts.m)

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

	for i := 0; i < 10; i++ {
		randomIndex := rand.Intn(len(teams))
		randTeam := teams[randomIndex]
		randBugTracker := bugTrackers[randomIndex]

		randomIndex = rand.Intn(len(users))
		randUser := users[randomIndex]

		mod := models.Module{
			Name: fmt.Sprintf("x/mod-%d", i),
			Team: randTeam,
			Repo: "https://github.com/cosmos/cosmos-sdk",
			Authors: []models.User{
				godUser,
				randUser,
			},
			Version: fmt.Sprintf("v1.0.%d", i),
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
		name     string
		query    string
		cursor   uint
		limit    int
		expected map[string]bool
	}{
		{"empty query", "", 0, 100, map[string]bool{}},
		{"no matching query", "no match", 0, 100, map[string]bool{}},
		{"matches one record", "x/mod-1", 0, 100, map[string]bool{"x/mod-1": true}},
		{
			"matches all records (page 1)", "module", 0, 5,
			map[string]bool{
				"x/mod-0": true,
				"x/mod-1": true,
				"x/mod-2": true,
				"x/mod-3": true,
				"x/mod-4": true,
			},
		},
		{
			"matches all records (page 2)", "module", 5, 5,
			map[string]bool{
				"x/mod-5": true,
				"x/mod-6": true,
				"x/mod-7": true,
				"x/mod-8": true,
				"x/mod-9": true,
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		mts.Run(tc.name, func() {
			// no matching query
			modules, err := models.SearchModules(mts.gormDB, tc.query, tc.cursor, tc.limit)
			mts.Require().NoError(err)
			mts.Require().Len(modules, len(tc.expected))

			for _, m := range modules {
				mts.Require().Contains(tc.expected, m.Name)
			}
		})
	}
}

func (mts *ModelsTestSuite) TestGetUserByID() {
	resetDB(mts.T(), mts.m)

	mod := models.Module{
		Name: "x/bank",
		Team: "cosmonauts",
		Repo: "https://github.com/cosmos/cosmos-sdk",
		Authors: []models.User{
			{Name: "foo", Email: models.NewNullString("foo@cosmonauts.com")},
		},
		Version: "v1.0.0",
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
	resetDB(mts.T(), mts.m)

	users, err := models.GetAllUsers(mts.gormDB, 0, 10)
	mts.Require().NoError(err)
	mts.Require().Empty(users)

	for i := 0; i < 25; i++ {
		mod := models.Module{
			Name: fmt.Sprintf("x/bank-%d", i),
			Team: "cosmonauts",
			Repo: "https://github.com/cosmos/cosmos-sdk",
			Authors: []models.User{
				{Name: fmt.Sprintf("foo-%d", i), Email: models.NewNullString(fmt.Sprintf("foo%d@cosmonauts.com", i))},
			},
			Version: "v1.0.0",
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

	// first page (full)
	users, err = models.GetAllUsers(mts.gormDB, 0, 10)
	mts.Require().NoError(err)
	mts.Require().Len(users, 10)

	cursor := users[len(users)-1].ID
	mts.Require().Equal(uint(10), cursor)

	// second page (full)
	users, err = models.GetAllUsers(mts.gormDB, cursor, 10)
	mts.Require().NoError(err)
	mts.Require().Len(users, 10)

	cursor = users[len(users)-1].ID
	mts.Require().Equal(uint(20), cursor)

	// third page (partially full)
	users, err = models.GetAllUsers(mts.gormDB, cursor, 10)
	mts.Require().NoError(err)
	mts.Require().Len(users, 5)

	cursor = users[len(users)-1].ID
	mts.Require().Equal(uint(25), cursor)
}

func (mts *ModelsTestSuite) TestGetAllKeywords() {
	resetDB(mts.T(), mts.m)

	keywords, err := models.GetAllKeywords(mts.gormDB, 0, 10)
	mts.Require().NoError(err)
	mts.Require().Empty(keywords)

	for i := 0; i < 25; i++ {
		mod := models.Module{
			Name: fmt.Sprintf("x/bank-%d", i),
			Team: "cosmonauts",
			Repo: "https://github.com/cosmos/cosmos-sdk",
			Authors: []models.User{
				{Name: "foo", Email: models.NewNullString("foo@cosmonauts.com")},
			},
			Version: "v1.0.0",
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

	// first page (full)
	keywords, err = models.GetAllKeywords(mts.gormDB, 0, 10)
	mts.Require().NoError(err)
	mts.Require().Len(keywords, 10)

	cursor := keywords[len(keywords)-1].ID
	mts.Require().Equal(uint(10), cursor)

	// second page (full)
	keywords, err = models.GetAllKeywords(mts.gormDB, cursor, 10)
	mts.Require().NoError(err)
	mts.Require().Len(keywords, 10)

	cursor = keywords[len(keywords)-1].ID
	mts.Require().Equal(uint(20), cursor)

	// third page (partially full)
	keywords, err = models.GetAllKeywords(mts.gormDB, cursor, 10)
	mts.Require().NoError(err)
	mts.Require().Len(keywords, 5)

	cursor = keywords[len(keywords)-1].ID
	mts.Require().Equal(uint(25), cursor)
}

func resetDB(t *testing.T, m *migrate.Migrate) {
	t.Helper()

	require.NoError(t, m.Force(1))
	require.NoError(t, m.Down())
	require.NoError(t, m.Up())
}
