package apiserver

import (
	"database/sql"
	"net/http"

	"github.com/ArdentK/bmstu-6sem-ppo/internal/app/store/mysqlstore"
	"github.com/ArdentK/bmstu-6sem-ppo/internal/app/store/sqlstore"
	"github.com/gorilla/sessions"

	_ "github.com/go-sql-driver/mysql"
)

func Start(config *Config) error {
	db, err := newDB(config.Database, config.DatabaseURL)
	if err != nil {
		return nil
	}
	defer db.Close()

	store := sqlstore.New(db)
	sessionStore := sessions.NewCookieStore([]byte(config.SessionKey))
	srv := newServer(store, sessionStore, "./templates/*")

	if config.Database == "mysql" {
		store2 := mysqlstore.New(db)
		srv2 := newServer(store2, sessionStore, "./templates/*")
		return http.ListenAndServe(config.BindAddr, srv2)
	}

	return http.ListenAndServe(config.BindAddr, srv)
}

func newDB(database, databaseURL string) (*sql.DB, error) {
	db, err := sql.Open(database, databaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
