package db

import (
	"context"
	"database/sql"
	"errors"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/stretchr/testify/require"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"readly/configs"
	"readly/testdata"
	"testing"
	"time"
)

var querier Querier

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

func TestMain(m *testing.M) {
	config, err := configs.Load(filepath.Join(configs.ProjectRoot(), "/configs/env"))
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	// Run migrations before tests
	if err := runMigrations(config.DBDriver, config.DBSource); err != nil {
		log.Fatal("migration failed:", err)
	}

	h := &Adapter{}
	_, q := h.Connect(config.DBDriver, config.DBSource)
	querier = q
	os.Exit(m.Run())
}

func createRandomUser(t *testing.T) User {
	password := testdata.RandomString(16)
	hashedPassword, err := testdata.HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)

	arg := CreateUserParams{
		Name:           testdata.RandomString(12),
		Email:          testdata.RandomEmail(),
		HashedPassword: hashedPassword,
	}
	user, err := querier.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	return user
}

func runMigrations(dbDriver string, dbSource string) error {
	// Connect to database
	db, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		return err
	}
	defer db.Close()

	// Create a db driver instance
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}

	// Create migrate instance
	migrationPath := filepath.Join(configs.ProjectRoot(), "db/migration")
	m, err := migrate.NewWithDatabaseInstance(
		"file://"+migrationPath,
		"postgres",
		driver,
	)
	if err != nil {
		return err
	}

	// Run migrations
	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	log.Println("Migrations completed successfully")
	return nil
}

func EqualDate(t *testing.T, a, b time.Time) {
	require.Equal(t, a.Year(), b.Year())
	require.Equal(t, a.Month(), b.Month())
	require.Equal(t, a.Day(), b.Day())
}
