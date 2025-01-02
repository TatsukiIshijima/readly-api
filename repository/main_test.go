package repository

import (
	"math/rand"
	"os"
	"readly/test"
	"testing"
	"time"
)

var store *Store
var repo BookRepository

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

func TestMain(m *testing.M) {
	test.Connect()
	store = NewStore(test.DB)
	repo = NewBookRepository(store)
	os.Exit(m.Run())
}
