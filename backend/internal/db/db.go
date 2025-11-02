package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func ConnectDB(path string) (*sql.DB, error) {
	return sql.Open("sqlite3", path)
}
