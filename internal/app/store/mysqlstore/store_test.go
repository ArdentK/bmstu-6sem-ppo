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
		databaseURL = "host=localhost port=5432 dbname=fencing_test sslmode=disable user=postgres password=postgres"
	}

	if database == "" {
		database = "postgres"
	}

	os.Exit(m.Run())
}
