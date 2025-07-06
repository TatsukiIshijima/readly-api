//go:build test

package repository

import (
	"math/rand"
	"os"
	sqlc "readly/db/sqlc"
	"testing"
	"time"
)

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

func TestMain(m *testing.M) {
	setupMain()
	os.Exit(m.Run())
}

var userRepo UserRepository

func setupMain() {
	fa := sqlc.FakeAdapter{}
	_, q := fa.Connect("", "")
	userRepo = NewUserRepository(q)
}
