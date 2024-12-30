package repository

import (
	"math/rand"
	"os"
	"readly/test"
	"testing"
	"time"
)

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

func TestMain(m *testing.M) {
	test.Connect()
	os.Exit(m.Run())
}
