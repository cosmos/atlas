package models_test

import (
	"context"
	"database/sql"
	"fmt"
	"math/rand"
	"os"
	"testing"

	"github.com/dhui/dktest"
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
	fmt.Println("DB:", connStr)
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

// SetupTestSuite executes before each individual test.
func (mts *ModelsTestSuite) SetupTestSuite() {
	resetDB(mts.T(), mts.m)
}

// TearDownSuite executes after all the suite's test have finished.
func (mts *ModelsTestSuite) TearDownSuite() {
	mts.T().Log("tearing down test suite")
	mts.Require().NoError(mts.db.Close())
}

func TestModelsTestSuite(t *testing.T) {
	suite.Run(t, new(ModelsTestSuite))
}

func (mts *ModelsTestSuite) TestModuleCreate() {
	mts.Require().True(true)
}

func TestModelz(t *testing.T) {
	t.SkipNow()

	dktest.Run(t, "postgres:11-alpine", dktest.Options{PortRequired: true, ReadyFunc: pgReady, Env: map[string]string{"POSTGRES_HOST_AUTH_METHOD": "trust"}},
		func(t *testing.T, c dktest.ContainerInfo) {
			ip, _, err := c.FirstPort()
			require.NoError(t, err)

			dsn := fmt.Sprintf("host=%s port=5432 dbname=postgres sslmode=disable", ip)

			db, err := sql.Open("postgres", dsn)
			require.NoError(t, err)
			defer db.Close()

			require.NoError(t, db.Ping())

			driver, err := migratepg.WithInstance(db, &migratepg.Config{})
			require.NoError(t, err)

			path := os.Getenv("ATLAS_MIGRATIONS_DIR")
			require.NotEmpty(t, path)

			m, err := migrate.NewWithDatabaseInstance(fmt.Sprintf("file:///%s", path), "postgres", driver)
			require.NoError(t, err)

			gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: gormlogger.Discard})
			require.NoError(t, err)

			t.Run("Module", func(t *testing.T) {
				testModuleCreate(t, m, gormDB)
				testGetModuleByID(t, m, gormDB)
				testGetAllModules(t, m, gormDB)
				testModuleUpdateBasic(t, m, gormDB)
				testModuleUpdateBugTracker(t, m, gormDB)
				testModuleUpdateKeywords(t, m, gormDB)
				testModuleUpdateAuthors(t, m, gormDB)
				testModuleUpdateOwners(t, m, gormDB)
				testModuleUpdateVersion(t, m, gormDB)
				testModuleSearch(t, m, gormDB)
			})

			t.Run("User", func(t *testing.T) {
				testGetUserByID(t, m, gormDB)
				testGetAllUsers(t, m, gormDB)
			})

			t.Run("Keyword", func(t *testing.T) {
				testGetAllKeywords(t, m, gormDB)
			})
		})
}

func pgReady(ctx context.Context, c dktest.ContainerInfo) bool {
	ip, _, err := c.FirstPort()
	if err != nil {
		return false
	}

	connStr := fmt.Sprintf("host=%s port=5432 dbname=postgres sslmode=disable", ip)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return false
	}
	defer db.Close()

	return db.PingContext(ctx) == nil
}

func resetDB(t *testing.T, m *migrate.Migrate) {
	t.Helper()

	require.NoError(t, m.Force(1))
	require.NoError(t, m.Down())
	require.NoError(t, m.Up())
}

func testModuleCreate(t *testing.T, m *migrate.Migrate, db *gorm.DB) {
	resetDB(t, m)

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

		t.Run(tc.name, func(t *testing.T) {
			result, err := tc.module.Upsert(db)
			if tc.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.module.Name, result.Name)
				require.Equal(t, tc.module.Team, result.Team)
				require.Equal(t, tc.module.Description, result.Description)
				require.Equal(t, tc.module.Homepage, result.Homepage)
				require.Equal(t, tc.module.Documentation, result.Documentation)
			}
		})
	}
}

func testModuleUpdateBasic(t *testing.T, m *migrate.Migrate, db *gorm.DB) {
	resetDB(t, m)

	t.Run("update basic", func(t *testing.T) {
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

		record, err := mod.Upsert(db)
		require.NoError(t, err)

		mod = models.Module{
			Name:          record.Name,
			Team:          record.Team,
			Description:   "new test description",
			Documentation: "https://github.com/cosmos/cosmos-sdk/x/bank/new_readme.md",
			Homepage:      "https://new.cosmos.network",
			Repo:          "https://github.com/cosmos/cosmos-sdk/new",
		}

		record, err = mod.Upsert(db)
		require.NoError(t, err)
		require.Equal(t, mod.Description, record.Description)
		require.Equal(t, mod.Documentation, record.Documentation)
		require.Equal(t, mod.Homepage, record.Homepage)
		require.Equal(t, mod.Repo, record.Repo)
	})
}

func testModuleUpdateBugTracker(t *testing.T, m *migrate.Migrate, db *gorm.DB) {
	resetDB(t, m)

	t.Run("update bug tracker", func(t *testing.T) {
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

		record, err := mod.Upsert(db)
		require.NoError(t, err)

		mod.BugTracker = models.BugTracker{
			URL:     models.NewNullString("https://cosmos.network/bugs"),
			Contact: models.NewNullString("bugs@cosmos.network"),
		}

		record, err = mod.Upsert(db)
		require.NoError(t, err)
		require.Equal(t, mod.BugTracker.URL, record.BugTracker.URL)
		require.Equal(t, mod.BugTracker.Contact, record.BugTracker.Contact)

		mod.BugTracker = models.BugTracker{}

		record, err = mod.Upsert(db)
		require.NoError(t, err)
		require.Equal(t, mod.BugTracker.URL, record.BugTracker.URL)
		require.Equal(t, mod.BugTracker.Contact, record.BugTracker.Contact)
	})
}

func testModuleUpdateKeywords(t *testing.T, m *migrate.Migrate, db *gorm.DB) {
	resetDB(t, m)

	t.Run("update keywords", func(t *testing.T) {
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

		record, err := mod.Upsert(db)
		require.NoError(t, err)

		mod.Keywords = []models.Keyword{
			{Name: "keyword1"}, {Name: "keyword2"}, {Name: "keyword3"},
		}

		record, err = mod.Upsert(db)
		require.NoError(t, err)
		require.Equal(t, mod.Keywords, record.Keywords)

		mod.Keywords = []models.Keyword{
			{Name: "keyword1"}, {Name: "keyword3"},
		}

		record, err = mod.Upsert(db)
		require.NoError(t, err)
		require.Equal(t, mod.Keywords, record.Keywords)

		mod.Keywords = []models.Keyword{}

		record, err = mod.Upsert(db)
		require.NoError(t, err)
		require.Equal(t, mod.Keywords, record.Keywords)
	})
}

func testModuleUpdateAuthors(t *testing.T, m *migrate.Migrate, db *gorm.DB) {
	resetDB(t, m)

	t.Run("update authors", func(t *testing.T) {
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

		record, err := mod.Upsert(db)
		require.NoError(t, err)

		mod.Authors = []models.User{
			{Name: "admin"}, {Name: "user1"}, {Name: "user2"},
		}

		record, err = mod.Upsert(db)
		require.NoError(t, err)
		require.Equal(t, mod.Authors, record.Authors)
		require.Equal(t, mod.Owners, record.Owners)

		mod.Authors = []models.User{
			{Name: "admin"}, {Name: "user2"},
		}

		record, err = mod.Upsert(db)
		require.NoError(t, err)
		require.Equal(t, mod.Authors, record.Authors)
		require.Equal(t, mod.Owners, record.Owners)

		mod.Authors = []models.User{}

		record, err = mod.Upsert(db)
		require.NoError(t, err)
		require.Equal(t, mod.Authors, record.Authors)
		require.Equal(t, mod.Owners, record.Owners)
	})
}

func testModuleUpdateOwners(t *testing.T, m *migrate.Migrate, db *gorm.DB) {
	resetDB(t, m)

	t.Run("update owners", func(t *testing.T) {
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

		record, err := mod.Upsert(db)
		require.NoError(t, err)

		mod.Owners = []models.User{
			{Name: "admin"}, {Name: "user1"}, {Name: "user2"},
		}

		record, err = mod.Upsert(db)
		require.NoError(t, err)
		require.Equal(t, mod.Authors, record.Authors)
		require.Equal(t, mod.Owners, record.Owners)

		mod.Owners = []models.User{
			{Name: "admin"}, {Name: "user2"},
		}

		record, err = mod.Upsert(db)
		require.NoError(t, err)
		require.Equal(t, mod.Authors, record.Authors)
		require.Equal(t, mod.Owners, record.Owners)

		mod.Owners = []models.User{}

		record, err = mod.Upsert(db)
		require.NoError(t, err)
		require.Equal(t, mod.Authors, record.Authors)
		require.Equal(t, mod.Owners, record.Owners)
	})
}

func testModuleUpdateVersion(t *testing.T, m *migrate.Migrate, db *gorm.DB) {
	resetDB(t, m)

	t.Run("update version", func(t *testing.T) {
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

		record, err := mod.Upsert(db)
		require.NoError(t, err)

		// update version
		mod.Version = "v1.0.1"

		record, err = mod.Upsert(db)
		require.NoError(t, err)
		require.Len(t, record.Versions, 2)

		latest, err := record.GetLatestVersion(db)
		require.NoError(t, err)
		require.Equal(t, mod.Version, latest.Version)

		// no version update
		mod.Version = "v1.0.1"

		record, err = mod.Upsert(db)
		require.NoError(t, err)
		require.Len(t, record.Versions, 2)

		latest, err = record.GetLatestVersion(db)
		require.NoError(t, err)
		require.Equal(t, mod.Version, latest.Version)

		// update version again
		mod.Version = "v2.0.0"

		record, err = mod.Upsert(db)
		require.NoError(t, err)
		require.Len(t, record.Versions, 3)

		latest, err = record.GetLatestVersion(db)
		require.NoError(t, err)
		require.Equal(t, mod.Version, latest.Version)
	})
}

func testModuleSearch(t *testing.T, m *migrate.Migrate, db *gorm.DB) {
	resetDB(t, m)

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

		_, err := mod.Upsert(db)
		require.NoError(t, err)
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

		t.Run(tc.name, func(t *testing.T) {
			// no matching query
			modules, err := models.SearchModules(db, tc.query, tc.cursor, tc.limit)
			require.NoError(t, err)
			require.Len(t, modules, len(tc.expected))

			for _, m := range modules {
				require.Contains(t, tc.expected, m.Name)
			}
		})
	}
}

func testGetModuleByID(t *testing.T, m *migrate.Migrate, db *gorm.DB) {
	resetDB(t, m)

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

	mod, err := mod.Upsert(db)
	require.NoError(t, err)

	t.Run("no module exists", func(t *testing.T) {
		result, err := models.GetModuleByID(db, mod.ID+1)
		require.Error(t, err)
		require.Equal(t, models.Module{}, result)
	})

	t.Run("module exists", func(t *testing.T) {
		result, err := models.GetModuleByID(db, mod.ID)
		require.NoError(t, err)
		require.Equal(t, mod.Name, result.Name)
		require.Equal(t, mod.Team, result.Team)
		require.Equal(t, mod.Description, result.Description)
		require.Equal(t, mod.Homepage, result.Homepage)
		require.Equal(t, mod.Documentation, result.Documentation)
	})
}

func testGetUserByID(t *testing.T, m *migrate.Migrate, db *gorm.DB) {
	resetDB(t, m)

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

	mod, err := mod.Upsert(db)
	require.NoError(t, err)

	t.Run("no user exists", func(t *testing.T) {
		result, err := models.GetUserByID(db, mod.Authors[0].ID+1)
		require.Error(t, err)
		require.Equal(t, models.User{}, result)
	})

	t.Run("user exists", func(t *testing.T) {
		result, err := models.GetUserByID(db, mod.Authors[0].ID)
		require.NoError(t, err)
		require.Equal(t, mod.Authors[0].Name, result.Name)
		require.Equal(t, mod.Authors[0].Email, result.Email)
	})
}

func testGetAllModules(t *testing.T, m *migrate.Migrate, db *gorm.DB) {
	resetDB(t, m)

	mods, err := models.GetAllModules(db, 0, 10)
	require.NoError(t, err)
	require.Empty(t, mods)

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

		_, err := mod.Upsert(db)
		require.NoError(t, err)
	}

	// first page (full)
	mods, err = models.GetAllModules(db, 0, 10)
	require.NoError(t, err)
	require.Len(t, mods, 10)

	cursor := mods[len(mods)-1].ID
	require.Equal(t, uint(10), cursor)

	// second page (full)
	mods, err = models.GetAllModules(db, cursor, 10)
	require.NoError(t, err)
	require.Len(t, mods, 10)

	cursor = mods[len(mods)-1].ID
	require.Equal(t, uint(20), cursor)

	// third page (partially full)
	mods, err = models.GetAllModules(db, cursor, 10)
	require.NoError(t, err)
	require.Len(t, mods, 5)

	cursor = mods[len(mods)-1].ID
	require.Equal(t, uint(25), cursor)
}

func testGetAllUsers(t *testing.T, m *migrate.Migrate, db *gorm.DB) {
	resetDB(t, m)

	users, err := models.GetAllUsers(db, 0, 10)
	require.NoError(t, err)
	require.Empty(t, users)

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

		_, err := mod.Upsert(db)
		require.NoError(t, err)
	}

	// first page (full)
	users, err = models.GetAllUsers(db, 0, 10)
	require.NoError(t, err)
	require.Len(t, users, 10)

	cursor := users[len(users)-1].ID
	require.Equal(t, uint(10), cursor)

	// second page (full)
	users, err = models.GetAllUsers(db, cursor, 10)
	require.NoError(t, err)
	require.Len(t, users, 10)

	cursor = users[len(users)-1].ID
	require.Equal(t, uint(20), cursor)

	// third page (partially full)
	users, err = models.GetAllUsers(db, cursor, 10)
	require.NoError(t, err)
	require.Len(t, users, 5)

	cursor = users[len(users)-1].ID
	require.Equal(t, uint(25), cursor)
}

func testGetAllKeywords(t *testing.T, m *migrate.Migrate, db *gorm.DB) {
	resetDB(t, m)

	keywords, err := models.GetAllKeywords(db, 0, 10)
	require.NoError(t, err)
	require.Empty(t, keywords)

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

		_, err := mod.Upsert(db)
		require.NoError(t, err)
	}

	// first page (full)
	keywords, err = models.GetAllKeywords(db, 0, 10)
	require.NoError(t, err)
	require.Len(t, keywords, 10)

	cursor := keywords[len(keywords)-1].ID
	require.Equal(t, uint(10), cursor)

	// second page (full)
	keywords, err = models.GetAllKeywords(db, cursor, 10)
	require.NoError(t, err)
	require.Len(t, keywords, 10)

	cursor = keywords[len(keywords)-1].ID
	require.Equal(t, uint(20), cursor)

	// third page (partially full)
	keywords, err = models.GetAllKeywords(db, cursor, 10)
	require.NoError(t, err)
	require.Len(t, keywords, 5)

	cursor = keywords[len(keywords)-1].ID
	require.Equal(t, uint(25), cursor)
}
