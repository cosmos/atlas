package v1

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
	"time"

	"github.com/dghubble/gologin/v2"
	"github.com/dghubble/gologin/v2/github"
	oauth2login "github.com/dghubble/gologin/v2/oauth2"
	"github.com/golang-migrate/migrate/v4"
	migratepg "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	gogithub "github.com/google/go-github/github"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/knadh/koanf"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"golang.org/x/oauth2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"

	"github.com/cosmos/atlas/server/httputil"
	"github.com/cosmos/atlas/server/models"
)

type RouterTestSuite struct {
	suite.Suite

	m      *migrate.Migrate
	db     *sql.DB
	mux    *mux.Router
	router *Router
}

// SetupSuite executes once before the suite's tests are executed.
func (rts *RouterTestSuite) SetupSuite() {
	migrationsPath := os.Getenv("ATLAS_MIGRATIONS_DIR")
	rts.Require().NotEmpty(migrationsPath)

	connStr := os.Getenv("ATLAS_TEST_DATABASE_URL")
	rts.Require().NotEmpty(connStr)

	db, err := sql.Open("postgres", connStr)
	rts.Require().NoError(err)
	rts.Require().NoError(db.Ping())

	driver, err := migratepg.WithInstance(db, &migratepg.Config{})
	rts.Require().NoError(err)

	m, err := migrate.NewWithDatabaseInstance(fmt.Sprintf("file:///%s", migrationsPath), "postgres", driver)
	rts.Require().NoError(err)

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
	rts.Require().NoError(err)

	sessionStore := sessions.NewCookieStore([]byte("service_test"), nil)
	sessionStore.Options.HttpOnly = true
	sessionStore.Options.Secure = false

	router, err := NewRouter(
		zerolog.New(ioutil.Discard).With().Timestamp().Logger(),
		koanf.New("."),
		gormDB,
		gologin.DebugOnlyCookieConfig,
		sessionStore,
		&oauth2.Config{},
	)
	rts.Require().NoError(err)

	mux := mux.NewRouter()
	router.Register(mux, V1APIPathPrefix)

	rts.m = m
	rts.db = db
	rts.router = router
	rts.mux = mux
}

// TearDownSuite executes after all the suite's test have finished.
func (mts *RouterTestSuite) TearDownSuite() {
	mts.Require().NoError(mts.db.Close())
}

func TestServiceTestSuite(t *testing.T) {
	suite.Run(t, new(RouterTestSuite))
}

func (rts *RouterTestSuite) executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	rts.mux.ServeHTTP(rr, req)

	return rr
}

func (rts *RouterTestSuite) authorizeRequest(req *http.Request, token, login string, id int64) *http.Request {
	rr := httptest.NewRecorder()

	ctx := oauth2login.WithToken(context.Background(), &oauth2.Token{AccessToken: token})
	ctx = github.WithUser(ctx, &gogithub.User{Login: &login, ID: &id})
	req = req.WithContext(ctx)

	rts.router.authorizeHandler().ServeHTTP(rr, req)
	rts.Require().Equal(http.StatusFound, rr.Code)

	return req
}

func (rts *RouterTestSuite) TestHealth() {
	rts.resetDB()

	req, err := http.NewRequest("GET", "/api/v1/health", nil)
	rts.Require().NoError(err)

	response := rts.executeRequest(req)
	rts.Require().Equal(http.StatusOK, response.Code)

	var health map[string]interface{}
	rts.Require().NoError(json.Unmarshal(response.Body.Bytes(), &health))
	rts.Require().Equal(health["status"], "ok")
}

func (rts *RouterTestSuite) TestSearchModules() {
	rts.resetDB()

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
			Version: models.ModuleVersion{Version: fmt.Sprintf("v1.0.%d", i)},
			Keywords: []models.Keyword{
				{Name: "module"},
				{Name: fmt.Sprintf("mod-keyword-%d", i+1)},
				{Name: fmt.Sprintf("mod-keyword-%d", (i+1)*3)},
			},
			BugTracker: randBugTracker,
		}

		_, err := mod.Upsert(rts.router.db)
		rts.Require().NoError(err)
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

		rts.Run(tc.name, func() {
			path := fmt.Sprintf("/api/v1/modules/search?cursor=%d&limit=%d&q=%s", tc.cursor, tc.limit, tc.query)
			req, err := http.NewRequest("GET", path, nil)
			rts.Require().NoError(err)

			response := rts.executeRequest(req)

			var pr httputil.PaginationResponse
			rts.Require().NoError(json.Unmarshal(response.Body.Bytes(), &pr))

			rts.Require().Equal(tc.cursor, pr.Cursor)
			rts.Require().Equal(tc.limit, pr.Limit)
			rts.Require().Equal(len(tc.expected), len(pr.Results.([]interface{})))

			for _, iFace := range pr.Results.([]interface{}) {
				m := iFace.(map[string]interface{})
				rts.Require().Contains(tc.expected, m["name"])
			}
		})
	}
}

func (rts *RouterTestSuite) TestGetAllModules() {
	rts.resetDB()

	path := fmt.Sprintf("/api/v1/modules?cursor=%d&limit=%d", 0, 10)
	req, err := http.NewRequest("GET", path, nil)
	rts.Require().NoError(err)

	response := rts.executeRequest(req)

	var pr httputil.PaginationResponse
	rts.Require().NoError(json.Unmarshal(response.Body.Bytes(), &pr))
	rts.Require().Empty(pr.Results)

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

		_, err := mod.Upsert(rts.router.db)
		rts.Require().NoError(err)
	}

	// first page (full)
	path = fmt.Sprintf("/api/v1/modules?cursor=%d&limit=%d", 0, 10)
	req, err = http.NewRequest("GET", path, nil)
	rts.Require().NoError(err)

	response = rts.executeRequest(req)
	rts.Require().NoError(json.Unmarshal(response.Body.Bytes(), &pr))
	rts.Require().Len(pr.Results, 10)

	mods := pr.Results.([]interface{})
	cursor := uint(mods[len(mods)-1].(map[string]interface{})["id"].(float64))
	rts.Require().Equal(uint(10), cursor)

	// second page (full)
	path = fmt.Sprintf("/api/v1/modules?cursor=%d&limit=%d", cursor, 10)
	req, err = http.NewRequest("GET", path, nil)
	rts.Require().NoError(err)

	response = rts.executeRequest(req)
	rts.Require().NoError(json.Unmarshal(response.Body.Bytes(), &pr))
	rts.Require().Len(pr.Results, 10)

	mods = pr.Results.([]interface{})
	cursor = uint(mods[len(mods)-1].(map[string]interface{})["id"].(float64))
	rts.Require().Equal(uint(20), cursor)

	// third page (partially full)
	path = fmt.Sprintf("/api/v1/modules?cursor=%d&limit=%d", cursor, 10)
	req, err = http.NewRequest("GET", path, nil)
	rts.Require().NoError(err)

	response = rts.executeRequest(req)
	rts.Require().NoError(json.Unmarshal(response.Body.Bytes(), &pr))
	rts.Require().Len(pr.Results, 5)

	mods = pr.Results.([]interface{})
	cursor = uint(mods[len(mods)-1].(map[string]interface{})["id"].(float64))
	rts.Require().Equal(uint(25), cursor)
}

func (rts *RouterTestSuite) TestGetModuleByID() {
	rts.resetDB()

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

	mod, err := mod.Upsert(rts.router.db)
	rts.Require().NoError(err)

	rts.Run("no module exists", func() {
		path := fmt.Sprintf("/api/v1/modules/%d", mod.ID+1)
		req, err := http.NewRequest("GET", path, nil)
		rts.Require().NoError(err)

		response := rts.executeRequest(req)

		var body map[string]interface{}
		rts.Require().NoError(json.Unmarshal(response.Body.Bytes(), &body))
		rts.Require().Equal(http.StatusNotFound, response.Code)
		rts.Require().NotEmpty(body["error"])
	})

	rts.Run("module exists", func() {
		path := fmt.Sprintf("/api/v1/modules/%d", mod.ID)
		req, err := http.NewRequest("GET", path, nil)
		rts.Require().NoError(err)

		response := rts.executeRequest(req)

		var body map[string]interface{}
		rts.Require().NoError(json.Unmarshal(response.Body.Bytes(), &body))
		rts.Require().Equal(http.StatusOK, response.Code)
		rts.Require().Equal(mod.Name, body["name"])
		rts.Require().Equal(mod.Team, body["team"])
		rts.Require().Equal(mod.Description, body["description"])
		rts.Require().Equal(mod.Homepage, body["homepage"])
		rts.Require().Equal(mod.Documentation, body["documentation"])
	})
}
func (rts *RouterTestSuite) TestGetModuleVersions() {
	rts.resetDB()

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

	mod, err := mod.Upsert(rts.router.db)
	rts.Require().NoError(err)

	rts.Run("no module exists", func() {
		path := fmt.Sprintf("/api/v1/modules/%d/versions", mod.ID+1)
		req, err := http.NewRequest("GET", path, nil)
		rts.Require().NoError(err)

		response := rts.executeRequest(req)

		var body map[string]interface{}
		rts.Require().NoError(json.Unmarshal(response.Body.Bytes(), &body))
		rts.Require().Equal(http.StatusNotFound, response.Code)
		rts.Require().NotEmpty(body["error"])
	})

	rts.Run("module exists", func() {
		path := fmt.Sprintf("/api/v1/modules/%d/versions", mod.ID)
		req, err := http.NewRequest("GET", path, nil)
		rts.Require().NoError(err)

		response := rts.executeRequest(req)

		var body []interface{}
		rts.Require().NoError(json.Unmarshal(response.Body.Bytes(), &body))
		rts.Require().Equal(http.StatusOK, response.Code)
		rts.Require().Len(body, 1)
		rts.Require().Equal("v1.0.0", body[0].(map[string]interface{})["version"])
	})
}

func (rts *RouterTestSuite) GetModuleAuthors() {
	rts.resetDB()

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

	mod, err := mod.Upsert(rts.router.db)
	rts.Require().NoError(err)

	rts.Run("no module exists", func() {
		path := fmt.Sprintf("/api/v1/modules/%d/authors", mod.ID+1)
		req, err := http.NewRequest("GET", path, nil)
		rts.Require().NoError(err)

		response := rts.executeRequest(req)

		var body map[string]interface{}
		rts.Require().NoError(json.Unmarshal(response.Body.Bytes(), &body))
		rts.Require().Equal(http.StatusNotFound, response.Code)
		rts.Require().NotEmpty(body["error"])
	})

	rts.Run("module exists", func() {
		path := fmt.Sprintf("/api/v1/modules/%d/authors", mod.ID)
		req, err := http.NewRequest("GET", path, nil)
		rts.Require().NoError(err)

		response := rts.executeRequest(req)

		var body []interface{}
		rts.Require().NoError(json.Unmarshal(response.Body.Bytes(), &body))
		rts.Require().Equal(http.StatusOK, response.Code)
		rts.Require().Len(body, 1)
		rts.Require().Equal("foo", body[0].(map[string]interface{})["name"])
	})
}

func (rts *RouterTestSuite) GetModuleKeywords() {
	rts.resetDB()

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

	mod, err := mod.Upsert(rts.router.db)
	rts.Require().NoError(err)

	rts.Run("no module exists", func() {
		path := fmt.Sprintf("/api/v1/modules/%d/keywords", mod.ID+1)
		req, err := http.NewRequest("GET", path, nil)
		rts.Require().NoError(err)

		response := rts.executeRequest(req)

		var body map[string]interface{}
		rts.Require().NoError(json.Unmarshal(response.Body.Bytes(), &body))
		rts.Require().Equal(http.StatusNotFound, response.Code)
		rts.Require().NotEmpty(body["error"])
	})

	rts.Run("module exists", func() {
		path := fmt.Sprintf("/api/v1/modules/%d/keywords", mod.ID)
		req, err := http.NewRequest("GET", path, nil)
		rts.Require().NoError(err)

		response := rts.executeRequest(req)

		var body []interface{}
		rts.Require().NoError(json.Unmarshal(response.Body.Bytes(), &body))
		rts.Require().Equal(http.StatusOK, response.Code)
		rts.Require().Len(body, 1)
		rts.Require().Equal("tokens", body[0].(map[string]interface{})["name"])
	})
}

func (rts *RouterTestSuite) GetUserByID() {
	rts.resetDB()

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

	mod, err := mod.Upsert(rts.router.db)
	rts.Require().NoError(err)

	rts.Run("no user exists", func() {
		path := fmt.Sprintf("/api/v1/users/%d", mod.Authors[0].ID+1)
		req, err := http.NewRequest("GET", path, nil)
		rts.Require().NoError(err)

		response := rts.executeRequest(req)

		var body map[string]interface{}
		rts.Require().NoError(json.Unmarshal(response.Body.Bytes(), &body))
		rts.Require().Equal(http.StatusNotFound, response.Code)
		rts.Require().NotEmpty(body["error"])
	})

	rts.Run("user exists", func() {
		path := fmt.Sprintf("/api/v1/users/%d", mod.Authors[0].ID)
		req, err := http.NewRequest("GET", path, nil)
		rts.Require().NoError(err)

		response := rts.executeRequest(req)

		var body map[string]interface{}
		rts.Require().NoError(json.Unmarshal(response.Body.Bytes(), &body))
		rts.Require().Equal(http.StatusOK, response.Code)
		rts.Require().Equal(mod.Authors[0].Name, body["name"])
		rts.Require().Equal(mod.Authors[0].Email, body["email"])
	})
}

func (rts *RouterTestSuite) TestGetUserModules() {
	rts.resetDB()

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

	mod, err := mod.Upsert(rts.router.db)
	rts.Require().NoError(err)

	rts.Run("no user exists", func() {
		path := fmt.Sprintf("/api/v1/users/%d/modules", mod.Authors[0].ID+1)
		req, err := http.NewRequest("GET", path, nil)
		rts.Require().NoError(err)

		response := rts.executeRequest(req)

		var body map[string]interface{}
		rts.Require().NoError(json.Unmarshal(response.Body.Bytes(), &body))
		rts.Require().Equal(http.StatusNotFound, response.Code)
		rts.Require().NotEmpty(body["error"])
	})

	rts.Run("user exists", func() {
		path := fmt.Sprintf("/api/v1/users/%d/modules", mod.Authors[0].ID)
		req, err := http.NewRequest("GET", path, nil)
		rts.Require().NoError(err)

		response := rts.executeRequest(req)

		var body []interface{}
		rts.Require().NoError(json.Unmarshal(response.Body.Bytes(), &body))
		rts.Require().Equal(http.StatusOK, response.Code)
		rts.Require().Len(body, 1)
		rts.Require().Equal(mod.Name, body[0].(map[string]interface{})["name"])
		rts.Require().Equal(mod.Team, body[0].(map[string]interface{})["team"])
	})
}

func (rts *RouterTestSuite) TestGetAllUsers() {
	rts.resetDB()

	path := fmt.Sprintf("/api/v1/users?cursor=%d&limit=%d", 0, 10)
	req, err := http.NewRequest("GET", path, nil)
	rts.Require().NoError(err)

	response := rts.executeRequest(req)

	var pr httputil.PaginationResponse
	rts.Require().NoError(json.Unmarshal(response.Body.Bytes(), &pr))
	rts.Require().Empty(pr.Results)

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

		_, err := mod.Upsert(rts.router.db)
		rts.Require().NoError(err)
	}

	// first page (full)
	path = fmt.Sprintf("/api/v1/users?cursor=%d&limit=%d", 0, 10)
	req, err = http.NewRequest("GET", path, nil)
	rts.Require().NoError(err)

	response = rts.executeRequest(req)
	rts.Require().NoError(json.Unmarshal(response.Body.Bytes(), &pr))
	rts.Require().Len(pr.Results, 10)

	users := pr.Results.([]interface{})
	cursor := uint(users[len(users)-1].(map[string]interface{})["id"].(float64))
	rts.Require().Equal(uint(10), cursor)

	// second page (full)
	path = fmt.Sprintf("/api/v1/users?cursor=%d&limit=%d", cursor, 10)
	req, err = http.NewRequest("GET", path, nil)
	rts.Require().NoError(err)

	response = rts.executeRequest(req)
	rts.Require().NoError(json.Unmarshal(response.Body.Bytes(), &pr))
	rts.Require().Len(pr.Results, 10)

	users = pr.Results.([]interface{})
	cursor = uint(users[len(users)-1].(map[string]interface{})["id"].(float64))
	rts.Require().Equal(uint(20), cursor)

	// third page (partially full)
	path = fmt.Sprintf("/api/v1/users?cursor=%d&limit=%d", cursor, 10)
	req, err = http.NewRequest("GET", path, nil)
	rts.Require().NoError(err)

	response = rts.executeRequest(req)
	rts.Require().NoError(json.Unmarshal(response.Body.Bytes(), &pr))
	rts.Require().Len(pr.Results, 5)

	users = pr.Results.([]interface{})
	cursor = uint(users[len(users)-1].(map[string]interface{})["id"].(float64))
	rts.Require().Equal(uint(25), cursor)
}

func (rts *RouterTestSuite) TestGetAllKeywords() {
	rts.resetDB()

	path := fmt.Sprintf("/api/v1/keywords?cursor=%d&limit=%d", 0, 10)
	req, err := http.NewRequest("GET", path, nil)
	rts.Require().NoError(err)

	response := rts.executeRequest(req)

	var pr httputil.PaginationResponse
	rts.Require().NoError(json.Unmarshal(response.Body.Bytes(), &pr))
	rts.Require().Empty(pr.Results)

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

		_, err := mod.Upsert(rts.router.db)
		rts.Require().NoError(err)
	}

	// first page (full)
	path = fmt.Sprintf("/api/v1/keywords?cursor=%d&limit=%d", 0, 10)
	req, err = http.NewRequest("GET", path, nil)
	rts.Require().NoError(err)

	response = rts.executeRequest(req)
	rts.Require().NoError(json.Unmarshal(response.Body.Bytes(), &pr))
	rts.Require().Len(pr.Results, 10)

	keywords := pr.Results.([]interface{})
	cursor := uint(keywords[len(keywords)-1].(map[string]interface{})["id"].(float64))
	rts.Require().Equal(uint(10), cursor)

	// second page (full)
	path = fmt.Sprintf("/api/v1/keywords?cursor=%d&limit=%d", cursor, 10)
	req, err = http.NewRequest("GET", path, nil)
	rts.Require().NoError(err)

	response = rts.executeRequest(req)
	rts.Require().NoError(json.Unmarshal(response.Body.Bytes(), &pr))
	rts.Require().Len(pr.Results, 10)

	keywords = pr.Results.([]interface{})
	cursor = uint(keywords[len(keywords)-1].(map[string]interface{})["id"].(float64))
	rts.Require().Equal(uint(20), cursor)

	// third page (partially full)
	path = fmt.Sprintf("/api/v1/keywords?cursor=%d&limit=%d", cursor, 10)
	req, err = http.NewRequest("GET", path, nil)
	rts.Require().NoError(err)

	response = rts.executeRequest(req)
	rts.Require().NoError(json.Unmarshal(response.Body.Bytes(), &pr))
	rts.Require().Len(pr.Results, 5)

	keywords = pr.Results.([]interface{})
	cursor = uint(keywords[len(keywords)-1].(map[string]interface{})["id"].(float64))
	rts.Require().Equal(uint(25), cursor)
}

func (rts *RouterTestSuite) TestUpsertModule() {
	rts.resetDB()

	req, err := http.NewRequest("GET", "/", nil)
	rts.Require().NoError(err)

	req = rts.authorizeRequest(req, "test_token", "test_user", 12345)

	upsertURL, err := url.Parse("/api/v1/modules")
	rts.Require().NoError(err)

	testCases := []struct {
		name string
		body map[string]interface{}
		code int
	}{
		{
			name: "invalid name",
			body: make(map[string]interface{}),
			code: http.StatusBadRequest,
		},
		{
			name: "invalid team",
			body: map[string]interface{}{
				"module": map[string]interface{}{
					"name": "x/bank",
				},
			},
			code: http.StatusBadRequest,
		},
		{
			name: "missing repo",
			body: map[string]interface{}{
				"module": map[string]interface{}{
					"name": "x/bank",
					"team": "cosmonauts",
				},
			},
			code: http.StatusBadRequest,
		},
		{
			name: "missing authors",
			body: map[string]interface{}{
				"module": map[string]interface{}{
					"name": "x/bank",
					"team": "cosmonauts",
					"repo": "https://github.com/cosmos/cosmos-sdk",
				},
			},
			code: http.StatusBadRequest,
		},
		{
			name: "missing version",
			body: map[string]interface{}{
				"module": map[string]interface{}{
					"name": "x/bank",
					"team": "cosmonauts",
					"repo": "https://github.com/cosmos/cosmos-sdk",
				},
				"authors": []map[string]interface{}{
					{
						"name": "foo", "email": "foo@email.com",
					},
				},
			},
			code: http.StatusBadRequest,
		},
		{
			name: "duplicate authors",
			body: map[string]interface{}{
				"module": map[string]interface{}{
					"name":     "x/bank",
					"team":     "cosmonauts",
					"repo":     "https://github.com/cosmos/cosmos-sdk",
					"keywords": []string{"tokens"},
				},
				"authors": []map[string]interface{}{
					{
						"name": "foo", "email": "foo@email.com",
					},
					{
						"name": "foo", "email": "foo@email.com",
					},
				},
				"version": map[string]interface{}{
					"version": "v1.0.0",
				},
				"bug_tracker": map[string]interface{}{
					"url":     "https://cosmonauts.com",
					"contact": "contact@cosmonauts.com",
				},
			},
			code: http.StatusBadRequest,
		},
		{
			name: "duplicate keywords",
			body: map[string]interface{}{
				"module": map[string]interface{}{
					"name":     "x/bank",
					"team":     "cosmonauts",
					"repo":     "https://github.com/cosmos/cosmos-sdk",
					"keywords": []string{"tokens", "tokens"},
				},
				"authors": []map[string]interface{}{
					{
						"name": "foo", "email": "foo@email.com",
					},
				},
				"version": map[string]interface{}{
					"version": "v1.0.0",
				},
				"bug_tracker": map[string]interface{}{
					"url":     "https://cosmonauts.com",
					"contact": "contact@cosmonauts.com",
				},
			},
			code: http.StatusBadRequest,
		},
		{
			name: "valid module",
			body: map[string]interface{}{
				"module": map[string]interface{}{
					"name":     "x/bank",
					"team":     "cosmonauts",
					"repo":     "https://github.com/cosmos/cosmos-sdk",
					"keywords": []string{"tokens"},
				},
				"authors": []map[string]interface{}{
					{
						"name": "foo", "email": "foo@email.com",
					},
				},
				"version": map[string]interface{}{
					"version": "v1.0.0",
				},
				"bug_tracker": map[string]interface{}{
					"url":     "https://cosmonauts.com",
					"contact": "contact@cosmonauts.com",
				},
			},
			code: http.StatusOK,
		},
	}

	for _, tc := range testCases {
		tc := tc

		rts.Run(tc.name, func() {
			bz, err := json.Marshal(tc.body)
			rts.Require().NoError(err)

			req.Method = httputil.MethodPUT
			req.URL = upsertURL
			req.Body = ioutil.NopCloser(bytes.NewBuffer(bz))
			req.ContentLength = int64(len(bz))

			rr := httptest.NewRecorder()
			rts.mux.ServeHTTP(rr, req)

			rts.Require().Equal(tc.code, rr.Code, rr.Body.String())
		})
	}
}

func (rts *RouterTestSuite) TestCreateModule_Unauthorized() {
	rts.resetDB()

	body := map[string]interface{}{
		"name": "x/bank",
		"team": "cosmonauts",
		"repo": "https://github.com/cosmos/cosmos-sdk",
		"authors": []map[string]interface{}{
			{
				"name": "foo", "email": "foo@email.com",
			},
		},
		"version":  "v1.0.0",
		"keywords": []string{"tokens"},
		"bug_tracker": map[string]interface{}{
			"url":     "https://cosmonauts.com",
			"contact": "contact@cosmonauts.com",
		},
	}

	bz, err := json.Marshal(body)
	rts.Require().NoError(err)

	req, err := http.NewRequest("PUT", "/api/v1/modules", bytes.NewBuffer(bz))
	rts.Require().NoError(err)

	response := rts.executeRequest(req)
	rts.Require().Equal(http.StatusUnauthorized, response.Code)
}

func (rts *RouterTestSuite) TestCreateModule_InvalidOwner() {
	rts.resetDB()

	req1, err := http.NewRequest("GET", "/", nil)
	rts.Require().NoError(err)

	req2, err := http.NewRequest("GET", "/", nil)
	rts.Require().NoError(err)

	req1 = rts.authorizeRequest(req1, "test_token1", "test_user1", 12345)
	req2 = rts.authorizeRequest(req2, "test_token2", "test_user2", 67899)

	upsertURL, err := url.Parse("/api/v1/modules")
	rts.Require().NoError(err)

	body := map[string]interface{}{
		"module": map[string]interface{}{
			"name":     "x/bank",
			"team":     "cosmonauts",
			"repo":     "https://github.com/cosmos/cosmos-sdk",
			"keywords": []string{"tokens"},
		},
		"authors": []map[string]interface{}{
			{
				"name": "foo", "email": "foo@email.com",
			},
		},
		"version": map[string]interface{}{
			"version": "v1.0.0",
		},
		"bug_tracker": map[string]interface{}{
			"url":     "https://cosmonauts.com",
			"contact": "contact@cosmonauts.com",
		},
	}

	// create module published by test_user1
	bz, err := json.Marshal(body)
	rts.Require().NoError(err)

	req1.Method = httputil.MethodPUT
	req1.URL = upsertURL
	req1.Body = ioutil.NopCloser(bytes.NewBuffer(bz))
	req1.ContentLength = int64(len(bz))

	rr := httptest.NewRecorder()
	rts.mux.ServeHTTP(rr, req1)
	rts.Require().Equal(http.StatusOK, rr.Code, rr.Body.String())

	// attempt to update module published by test_user2
	req2.Method = httputil.MethodPUT
	req2.URL = upsertURL
	req2.Body = ioutil.NopCloser(bytes.NewBuffer(bz))
	req2.ContentLength = int64(len(bz))

	rr = httptest.NewRecorder()
	rts.mux.ServeHTTP(rr, req2)
	rts.Require().Equal(http.StatusBadRequest, rr.Code, rr.Body.String())
}

func (rts *RouterTestSuite) TestCreateUserToken() {
	rts.resetDB()

	unAuthReq, err := http.NewRequest(httputil.MethodPUT, "/api/v1/user/tokens", nil)
	rts.Require().NoError(err)

	// unauthenticated
	rr := httptest.NewRecorder()
	rts.mux.ServeHTTP(rr, unAuthReq)
	rts.Require().Equal(http.StatusUnauthorized, rr.Code, rr.Body.String())

	// authenticated
	req, err := http.NewRequest(httputil.MethodGET, "/", nil)
	rts.Require().NoError(err)

	req = rts.authorizeRequest(req, "test_token1", "test_user1", 123456)
	req.Method = httputil.MethodPUT
	req.URL = unAuthReq.URL

	for i := int64(0); i < MaxTokens; i++ {
		rr = httptest.NewRecorder()
		rts.mux.ServeHTTP(rr, req)
		rts.Require().Equal(http.StatusOK, rr.Code, rr.Body.String())

		var ut map[string]interface{}
		rts.Require().NoError(json.Unmarshal(rr.Body.Bytes(), &ut), rr.Body.String())
		rts.Require().NotEmpty(ut["token"])
		rts.Require().Equal(1, int(ut["user_id"].(float64)))
	}

	// max tokens reached
	rr = httptest.NewRecorder()
	rts.mux.ServeHTTP(rr, req)
	rts.Require().Equal(http.StatusBadRequest, rr.Code, rr.Body.String())
}

func (rts *RouterTestSuite) TestGetUserTokens() {
	rts.resetDB()

	unAuthReq, err := http.NewRequest(httputil.MethodGET, "/api/v1/user/tokens", nil)
	rts.Require().NoError(err)

	// unauthenticated
	rr := httptest.NewRecorder()
	rts.mux.ServeHTTP(rr, unAuthReq)
	rts.Require().Equal(http.StatusUnauthorized, rr.Code, rr.Body.String())

	// authenticated
	req, err := http.NewRequest(httputil.MethodGET, "/", nil)
	rts.Require().NoError(err)

	req = rts.authorizeRequest(req, "test_token1", "test_user1", 123456)
	req.Method = httputil.MethodPUT
	req.URL = unAuthReq.URL

	for i := 0; i < 25; i++ {
		rr = httptest.NewRecorder()
		rts.mux.ServeHTTP(rr, req)
		rts.Require().Equal(http.StatusOK, rr.Code, rr.Body.String())

		var ut map[string]interface{}
		rts.Require().NoError(json.Unmarshal(rr.Body.Bytes(), &ut))
		rts.Require().NotEmpty(ut["token"])
		rts.Require().Equal(1, int(ut["user_id"].(float64)))
	}

	req.Method = httputil.MethodGET

	rr = httptest.NewRecorder()
	rts.mux.ServeHTTP(rr, req)
	rts.Require().Equal(http.StatusOK, rr.Code, rr.Body.String())

	var tokens []map[string]interface{}
	rts.Require().NoError(json.Unmarshal(rr.Body.Bytes(), &tokens))
	rts.Require().Len(tokens, 25)
}

func (rts *RouterTestSuite) TestRevokeUserToken() {
	rts.resetDB()

	unAuthReq, err := http.NewRequest(httputil.MethodDELETE, "/api/v1/user/tokens/1", nil)
	rts.Require().NoError(err)

	// unauthenticated
	rr := httptest.NewRecorder()
	rts.mux.ServeHTTP(rr, unAuthReq)
	rts.Require().Equal(http.StatusUnauthorized, rr.Code, rr.Body.String())

	// authenticated
	req, err := http.NewRequest(httputil.MethodGET, "/", nil)
	rts.Require().NoError(err)

	createURL, err := url.Parse("/api/v1/user/tokens")
	rts.Require().NoError(err)

	req = rts.authorizeRequest(req, "test_token1", "test_user1", 123456)
	req.Method = httputil.MethodPUT
	req.URL = createURL

	for i := 0; i < 25; i++ {
		rr = httptest.NewRecorder()
		rts.mux.ServeHTTP(rr, req)
		rts.Require().Equal(http.StatusOK, rr.Code, rr.Body.String())

		var ut map[string]interface{}
		rts.Require().NoError(json.Unmarshal(rr.Body.Bytes(), &ut))
		rts.Require().NotEmpty(ut["token"])
		rts.Require().Equal(1, int(ut["user_id"].(float64)))
	}

	req.Method = httputil.MethodDELETE
	req.URL = unAuthReq.URL

	rr = httptest.NewRecorder()
	rts.mux.ServeHTTP(rr, req)
	rts.Require().Equal(http.StatusOK, rr.Code, rr.Body.String())

	var ut map[string]interface{}
	rts.Require().NoError(json.Unmarshal(rr.Body.Bytes(), &ut))
	rts.Require().NotEmpty(ut["token"])
	rts.Require().True(ut["revoked"].(bool))

	// attempt to revoke an non-existant token
	revokeURL, err := url.Parse("/api/v1/user/tokens/100")
	rts.Require().NoError(err)

	req.URL = revokeURL

	rr = httptest.NewRecorder()
	rts.mux.ServeHTTP(rr, req)
	rts.Require().Equal(http.StatusNotFound, rr.Code, rr.Body.String())
}

func (rts *RouterTestSuite) resetDB() {
	rts.T().Helper()

	require.NoError(rts.T(), rts.m.Force(1))
	require.NoError(rts.T(), rts.m.Down())
	require.NoError(rts.T(), rts.m.Up())
}
