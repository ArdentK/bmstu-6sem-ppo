package mysqlstore_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
	database    string
)

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")
	database = os.Getenv("DATABASE")
	if databaseURL == "" {
		databaseURL = "user:Password1973!@/fencing_test"
	}

	if database == "" {
		database = "mysql"
	}

	os.Exit(m.Run())
}
