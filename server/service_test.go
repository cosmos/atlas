package server

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/dghubble/gologin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-migrate/migrate/v4"
	migratepg "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"golang.org/x/oauth2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"

	"github.com/cosmos/atlas/server/models"
)

type ServiceTestSuite struct {
	suite.Suite

	m       *migrate.Migrate
	db      *sql.DB
	gormDB  *gorm.DB
	service *Service
}

// SetupSuite executes once before the suite's tests are executed.
func (sts *ServiceTestSuite) SetupSuite() {
	migrationsPath := os.Getenv("ATLAS_MIGRATIONS_DIR")
	sts.Require().NotEmpty(migrationsPath)

	connStr := os.Getenv("ATLAS_TEST_DATABASE_URL")
	sts.Require().NotEmpty(connStr)

	db, err := sql.Open("postgres", connStr)
	sts.Require().NoError(err)
	sts.Require().NoError(db.Ping())

	driver, err := migratepg.WithInstance(db, &migratepg.Config{})
	sts.Require().NoError(err)

	m, err := migrate.NewWithDatabaseInstance(fmt.Sprintf("file:///%s", migrationsPath), "postgres", driver)
	sts.Require().NoError(err)

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
	sts.Require().NoError(err)

	sessionStore := sessions.NewCookieStore([]byte("service_test"), nil)
	sessionStore.Options.HttpOnly = true
	sessionStore.Options.Secure = false

	service := &Service{
		logger:       zerolog.New(ioutil.Discard).With().Timestamp().Logger(),
		db:           gormDB,
		validate:     validator.New(),
		router:       mux.NewRouter(),
		sessionStore: sessionStore,
		cookieCfg:    gologin.DebugOnlyCookieConfig,
		oauth2Cfg:    &oauth2.Config{},
	}

	service.registerV1Routes()

	sts.m = m
	sts.db = db
	sts.gormDB = gormDB
	sts.service = service
}

// TearDownSuite executes after all the suite's test have finished.
func (mts *ServiceTestSuite) TearDownSuite() {
	mts.Require().NoError(mts.db.Close())
}

func TestServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ServiceTestSuite))
}

func (sts *ServiceTestSuite) executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	sts.service.router.ServeHTTP(rr, req)

	return rr
}

func (sts *ServiceTestSuite) TestSearchModules() {
	resetDB(sts.T(), sts.m)

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

		_, err := mod.Upsert(sts.gormDB)
		sts.Require().NoError(err)
	}

	testCases := []struct {
		name     string
		query    string
		cursor   uint
		limit    int
		status   int
		expected map[string]bool
	}{
		{"empty query", "", 0, 100, 200, map[string]bool{}},
		{"no matching query", "no match", 0, 100, 200, map[string]bool{}},
		{"matches one record", "x/mod-1", 0, 100, 200, map[string]bool{"x/mod-1": true}},
		{
			"matches all records (page 1)", "module", 0, 5, 200,
			map[string]bool{
				"x/mod-0": true,
				"x/mod-1": true,
				"x/mod-2": true,
				"x/mod-3": true,
				"x/mod-4": true,
			},
		},
		{
			"matches all records (page 2)", "module", 5, 5, 200,
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

		sts.Run(tc.name, func() {
			path := fmt.Sprintf("/api/v1/modules/search?cursor=%d&limit=%d&q=%s", tc.cursor, tc.limit, tc.query)
			req, err := http.NewRequest("GET", path, nil)
			sts.Require().NoError(err)

			response := sts.executeRequest(req)

			var pr PaginationResponse
			sts.Require().NoError(json.Unmarshal(response.Body.Bytes(), &pr))

			sts.Require().Equal(tc.cursor, pr.Cursor)
			sts.Require().Equal(tc.limit, pr.Limit)
			sts.Require().Equal(len(tc.expected), len(pr.Results.([]interface{})))

			for _, iFace := range pr.Results.([]interface{}) {
				m := iFace.(map[string]interface{})
				sts.Require().Contains(tc.expected, m["name"])
			}
		})
	}
}

func resetDB(t *testing.T, m *migrate.Migrate) {
	t.Helper()

	require.NoError(t, m.Force(1))
	require.NoError(t, m.Down())
	require.NoError(t, m.Up())
}
