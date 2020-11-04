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
				Version: models.ModuleVersion{Version: "v1.0.0"},
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
	mts.resetDB()

	mod := models.Module{
		Name: "x/bank",
		Team: "cosmonauts",
		Repo: "https://github.com/cosmos/cosmos-sdk",
		Authors: []models.User{
			{Name: "foo", Email: models.NewNullString("foo@cosmonauts.com")},
		},
		Version: models.ModuleVersion{Version: "v1.0.0"},
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
			Repo: "https://github.com/cosmos/cosmos-sdk",
			Authors: []models.User{
				{Name: "foo", Email: models.NewNullString("foo@cosmonauts.com")},
			},
			Version: models.ModuleVersion{Version: "v1.0.0"},
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

func (mts *ModelsTestSuite) TestModuleUpdateBasic() {
	mts.resetDB()

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
		Version:    models.ModuleVersion{Version: "v1.0.0"},
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
	mts.resetDB()

	mod := models.Module{
		Name: "x/bank",
		Team: "cosmonauts",
		Repo: "https://github.com/cosmos/cosmos-sdk",
		Authors: []models.User{
			{Name: "admin"},
		},
		Version:    models.ModuleVersion{Version: "v1.0.0"},
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
		Repo: "https://github.com/cosmos/cosmos-sdk",
		Authors: []models.User{
			{Name: "admin"},
		},
		Version:    models.ModuleVersion{Version: "v1.0.0"},
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
		Repo: "https://github.com/cosmos/cosmos-sdk",
		Authors: []models.User{
			{Name: "admin"},
		},
		Owners: []models.User{
			{Name: "admin"},
		},
		Version:    models.ModuleVersion{Version: "v1.0.0"},
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
		Repo: "https://github.com/cosmos/cosmos-sdk",
		Authors: []models.User{
			{Name: "admin"},
		},
		Owners: []models.User{
			{Name: "admin"},
		},
		Version:    models.ModuleVersion{Version: "v1.0.0"},
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
		Repo: "https://github.com/cosmos/cosmos-sdk",
		Authors: []models.User{
			{Name: "admin"},
		},
		Owners: []models.User{
			{Name: "admin"},
		},
		Version:    models.ModuleVersion{Version: "v1.0.0"},
		BugTracker: models.BugTracker{},
	}

	record, err := mod.Upsert(mts.gormDB)
	mts.Require().NoError(err)

	// update version
	mod.Version = models.ModuleVersion{Version: "v1.0.1"}

	record, err = mod.Upsert(mts.gormDB)
	mts.Require().NoError(err)
	mts.Require().Len(record.Versions, 2)

	latest, err := record.GetLatestVersion(mts.gormDB)
	mts.Require().NoError(err)
	mts.Require().Equal(mod.Version.Version, latest.Version)
	mts.Require().Equal(mod.Version.SDKCompat, latest.SDKCompat)

	// no version update
	mod.Version = models.ModuleVersion{Version: "v1.0.1"}

	record, err = mod.Upsert(mts.gormDB)
	mts.Require().NoError(err)
	mts.Require().Len(record.Versions, 2)

	latest, err = record.GetLatestVersion(mts.gormDB)
	mts.Require().NoError(err)
	mts.Require().Equal(mod.Version.Version, latest.Version)
	mts.Require().Equal(mod.Version.SDKCompat, latest.SDKCompat)

	// update version again
	mod.Version = models.ModuleVersion{Version: "v2.0.0"}

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
			Repo: "https://github.com/cosmos/cosmos-sdk",
			Authors: []models.User{
				godUser,
				randUser,
			},
			Version: models.ModuleVersion{Version: fmt.Sprintf("v1.0.%d", i)},
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
		Repo: "https://github.com/cosmos/cosmos-sdk",
		Owners: []models.User{
			{Name: "foo", Email: models.NewNullString("foo@cosmonauts.com")},
		},
		Authors: []models.User{
			{Name: "foo", Email: models.NewNullString("foo@cosmonauts.com")},
		},
		Version: models.ModuleVersion{Version: "v1.0.0"},
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
		Repo: "https://github.com/cosmos/cosmos-sdk",
		Authors: []models.User{
			{Name: "foo", Email: models.NewNullString("foo@cosmonauts.com")},
		},
		Version: models.ModuleVersion{Version: "v1.0.0"},
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
			Repo: "https://github.com/cosmos/cosmos-sdk",
			Authors: []models.User{
				{Name: fmt.Sprintf("foo-%d", i), Email: models.NewNullString(fmt.Sprintf("foo%d@cosmonauts.com", i))},
			},
			Version: models.ModuleVersion{Version: "v1.0.0"},
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
			Repo: "https://github.com/cosmos/cosmos-sdk",
			Authors: []models.User{
				{Name: "foo", Email: models.NewNullString("foo@cosmonauts.com")},
			},
			Version: models.ModuleVersion{Version: "v1.0.0"},
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
		Repo: "https://github.com/cosmos/cosmos-sdk",
		Owners: []models.User{
			{Name: "foo", Email: models.NewNullString("foo@cosmonauts.com")},
		},
		Authors: []models.User{
			{Name: "foo", Email: models.NewNullString("foo@cosmonauts.com")},
		},
		Version: models.ModuleVersion{Version: "v1.0.0"},
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
		Repo: "https://github.com/cosmos/cosmos-sdk",
		Owners: []models.User{
			{Name: "foo", Email: models.NewNullString("foo@cosmonauts.com")},
		},
		Authors: []models.User{
			{Name: "foo", Email: models.NewNullString("foo@cosmonauts.com")},
		},
		Version: models.ModuleVersion{Version: "v1.0.0"},
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

	uec, err := models.UserEmailConfirmation{UserID: mod.Owners[0].ID}.Upsert(mts.gormDB)
	token1 := uec.Token
	mts.Require().NoError(err)
	mts.Require().NotEqual(uuid.UUID{}, token1)

	uec, err = models.UserEmailConfirmation{UserID: mod.Owners[0].ID}.Upsert(mts.gormDB)
	token2 := uec.Token
	mts.Require().NoError(err)
	mts.Require().NotEqual(uuid.UUID{}, token2)
	mts.Require().NotEqual(token1, token2)

	var count int64
	mts.Require().NoError(mts.gormDB.Table("user_email_confirmations").Count(&count).Error)
	mts.Require().Equal(int64(1), count)
}

func (mts *ModelsTestSuite) TestQueryUserEmailConfirmation() {
	mts.resetDB()

	mod := models.Module{
		Name: "x/bank",
		Team: "cosmonauts",
		Repo: "https://github.com/cosmos/cosmos-sdk",
		Owners: []models.User{
			{Name: "foo", Email: models.NewNullString("foo@cosmonauts.com")},
		},
		Authors: []models.User{
			{Name: "foo", Email: models.NewNullString("foo@cosmonauts.com")},
		},
		Version: models.ModuleVersion{Version: "v1.0.0"},
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

	uec, err = models.UserEmailConfirmation{UserID: mod.Owners[0].ID}.Upsert(mts.gormDB)
	token := uec.Token
	mts.Require().NoError(err)
	mts.Require().NotEqual(uuid.UUID{}, token)

	uec, err = models.QueryUserEmailConfirmation(mts.gormDB, query)
	mts.Require().NoError(err)
	mts.Require().Equal(mod.Owners[0].ID, uec.UserID)
	mts.Require().Equal(token, uec.Token)
}

func (mts *ModelsTestSuite) TestUser_ConfirmEmail() {
	mts.resetDB()

	mod := models.Module{
		Name: "x/bank",
		Team: "cosmonauts",
		Repo: "https://github.com/cosmos/cosmos-sdk",
		Owners: []models.User{
			{Name: "foo", Email: models.NewNullString("foo@cosmonauts.com")},
		},
		Authors: []models.User{
			{Name: "foo", Email: models.NewNullString("foo@cosmonauts.com")},
		},
		Version: models.ModuleVersion{Version: "v1.0.0"},
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

	uec, err := models.UserEmailConfirmation{UserID: mod.Owners[0].ID}.Upsert(mts.gormDB)
	token := uec.Token
	mts.Require().NoError(err)
	mts.Require().NotEqual(uuid.UUID{}, token)

	user, err = user.ConfirmEmail(mts.gormDB)
	mts.Require().NoError(err)
	mts.Require().True(user.EmailConfirmed)

	query := map[string]interface{}{"user_id": mod.Owners[0].ID}
	uec, err = models.QueryUserEmailConfirmation(mts.gormDB, query)
	mts.Require().Error(err)
	mts.Require().Empty(uec)
}

func (mts *ModelsTestSuite) resetDB() {
	mts.T().Helper()

	if err := mts.m.Down(); err != nil {
		require.Equal(mts.T(), migrate.ErrNoChange, err)
	}
	if err := mts.m.Up(); err != nil {
		require.Equal(mts.T(), migrate.ErrNoChange, err)
	}
}
