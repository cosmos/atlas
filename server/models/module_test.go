package models_test

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/cosmos/atlas/server/models"
	"github.com/dhui/dktest"
	"github.com/golang-migrate/migrate/v4"
	migratepg "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestModule(t *testing.T) {
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

			gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
			require.NoError(t, err)

			testModuleCreate(t, m, gormDB)
			testModuleUpdate(t, m, gormDB)
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
					{Name: "foo", Email: "foo@email.com"},
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
					{Name: "foo", Email: "foo@cosmonauts.com"},
				},
				Version: "v1.0.0",
				Keywords: []models.Keyword{
					{Name: "tokens"},
				},
				BugTracker: models.BugTracker{URL: "cosmonauts.com", Contact: "contact@cosmonauts.com"},
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

func testModuleUpdate(t *testing.T, m *migrate.Migrate, db *gorm.DB) {
	resetDB(t, m)

	mod := models.Module{
		Name: "x/bank",
		Team: "cosmonauts",
		Repo: "https://github.com/cosmos/cosmos-sdk",
		Authors: []models.User{
			{Name: "foo", Email: "foo@cosmonauts.com"},
		},
		Version: "v1.0.0",
		Keywords: []models.Keyword{
			{Name: "tokens"},
		},
		BugTracker: models.BugTracker{URL: "cosmonauts.com", Contact: "contact@cosmonauts.com"},
	}

	mod, err := mod.Upsert(db)
	require.NoError(t, err)

	t.Run("update module add author", func(t *testing.T) {
		mod.Authors = append(mod.Authors, models.User{Name: "bar", Email: "bar@cosmonauts.com"})
		result, err := mod.Upsert(db)
		require.NoError(t, err)
		require.Len(t, result.Authors, 2)
	})

	t.Run("update module remove author", func(t *testing.T) {
		mod.Authors = mod.Authors[:1]
		result, err := mod.Upsert(db)
		require.NoError(t, err)
		require.Len(t, result.Authors, 1)
	})

	t.Run("update module add keyword", func(t *testing.T) {
		mod.Keywords = append(mod.Keywords, models.Keyword{Name: "banking"})
		result, err := mod.Upsert(db)
		require.NoError(t, err)
		require.Len(t, result.Keywords, 2)
	})

	t.Run("update module remove keyword", func(t *testing.T) {
		mod.Keywords = mod.Keywords[:1]
		result, err := mod.Upsert(db)
		require.NoError(t, err)
		require.Len(t, result.Keywords, 1)
	})

	t.Run("update module bugtracker", func(t *testing.T) {
		mod.BugTracker = models.BugTracker{URL: "newcosmonauts.com"}
		result, err := mod.Upsert(db)
		require.NoError(t, err)
		require.Equal(t, "newcosmonauts.com", result.BugTracker.URL)
	})

	t.Run("update module version", func(t *testing.T) {
		mod.Version = "v1.0.1"
		result, err := mod.Upsert(db)
		require.NoError(t, err)
		require.Len(t, result.Versions, 2)
	})
}
