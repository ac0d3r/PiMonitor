package database

import (
	"errors"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	if err := Init("pi.sqlite3"); err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			panic(err)
		}
	}
	db = db.Debug()
	m.Run()
}
