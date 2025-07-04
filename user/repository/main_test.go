package repository

import (
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"readly/configs"
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

var config configs.Config
var querier sqlc.Querier
var userRepo UserRepository

func setupMain() {
	c, err := configs.Load(filepath.Join(configs.ProjectRoot(), "/configs/env"))
	if err != nil {
		log.Fatalf("cannot load config: %v", err)
	}
	config = c

	a := &sqlc.Adapter{}
	_, q := a.Connect(c.DBDriver, c.DBSource)
	querier = q

	userRepo = NewUserRepository(q)
}
