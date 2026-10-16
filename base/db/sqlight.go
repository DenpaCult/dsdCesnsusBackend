package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type sqliteDB struct {
	db *sql.DB
}

func NewSqliteDB() *sqliteDB {
	return &sqliteDB{nil}
}


func (db *sqliteDB) Init(address string) error {
	var err error
	db.db, err = sql.Open("sqlite3", address)
	if err != nil {
		return err
	}

	return nil
}

func (db *sqliteDB) Close() {
	if db.db != nil {
		db.db.Close()
	}
}

func (db *sqliteDB) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return db.db.Query(query, args...)
}

func (db *sqliteDB) QueryRow(query string, args ...interface{}) *sql.Row {
	return db.db.QueryRow(query, args...)
}

func (db *sqliteDB) Exec(query string, args ...interface{}) (sql.Result, error) {
	return db.db.Exec(query, args...)
}

func (db *sqliteDB) BeginTX () (*sql.Tx, error) {
	return db.db.Begin()
}


