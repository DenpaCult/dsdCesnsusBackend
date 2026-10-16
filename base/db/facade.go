package db

import (
	"database/sql"
	appconfig "dsdCensus/base/config"
	"log"
)

var (
	sqlite *sqliteDB
)

type DAO interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
	BeginTX() (*sql.Tx, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

func Init(config *appconfig.Config) *sqliteDB {
	sqlite = NewSqliteDB()

	sqlCfg := config.DataBase.Sqlite
	err := sqlite.Init(sqlCfg.Address)
	if err != nil {
		log.Panicf("Mysql init failed: %s", err.Error())
	}
	return sqlite
}

func Close() {
	sqlite.Close()
}

func SqliteDao() DAO {
	return sqlite
}
