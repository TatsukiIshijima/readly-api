package sqlc_test

import (
	"log"
	"math/rand"
	"os"
	"readly/env"
	"readly/test"
	"testing"
	"time"
)

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

func TestMain(m *testing.M) {
	config, err := env.Load("../../env")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	test.Connect(config.DBDriver, config.DBSource)
	os.Exit(m.Run())
}
