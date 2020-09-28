package db_test

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/dhui/dktest"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

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

func TestMigration(t *testing.T) {
	t.SkipNow()

	dktest.Run(t, "postgres:11-alpine", dktest.Options{PortRequired: true, ReadyFunc: pgReady, Env: map[string]string{"POSTGRES_HOST_AUTH_METHOD": "trust"}},
		func(t *testing.T, c dktest.ContainerInfo) {
			ip, _, err := c.FirstPort()
			require.NoError(t, err)

			connStr := fmt.Sprintf("host=%s port=5432 dbname=postgres sslmode=disable", ip)
			db, err := sql.Open("postgres", connStr)
			require.NoError(t, err)
			defer db.Close()

			require.NoError(t, db.Ping())

			driver, err := postgres.WithInstance(db, &postgres.Config{})
			require.NoError(t, err)

			path, err := os.Getwd()
			require.NoError(t, err)

			m, err := migrate.NewWithDatabaseInstance(fmt.Sprintf("file:///%s/migrations", path), "postgres", driver)
			require.NoError(t, err)

			require.NoError(t, m.Force(1))
			require.NoError(t, m.Down())
			require.NoError(t, m.Up())
		})
}
