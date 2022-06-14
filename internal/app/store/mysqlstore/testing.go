package mysqlstore

import (
	"database/sql"
	"fmt"
	"strings"
	"testing"
)

func TestDB(t *testing.T, database, databaseURL string) (*sql.DB, func(...string)) {
	t.Helper()

	db, err := sql.Open(database, databaseURL)
	if err != nil {
		t.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		t.Fatal(err)
	}

	return db, func(tables ...string) {
		if len(tables) > 0 {
			db.Exec(fmt.Sprintf("TRUNCATE table %s", strings.Join(tables, ", ")))
		}

		db.Close()
	}
}
