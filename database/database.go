package database

import (
	"database/sql"
	"fmt"

	"github.com/jdxj/video/config"
	"github.com/jdxj/video/logger"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dsnFormat = "%s:%s@tcp(%s)/%s?loc=Local&parseTime=true"
)

var (
	db *sql.DB
)

func Init() error {
	dbCfg := config.DB
	dsn := fmt.Sprintf(dsnFormat, dbCfg.User, dbCfg.Pass, dbCfg.Host, dbCfg.Base)

	mysql, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	if err := mysql.Ping(); err != nil {
		return err
	}
	db = mysql
	return nil
}

func Query(query string, args ...interface{}) (*sql.Rows, error) {
	return db.Query(query, args...)
}

func QueryRow(query string, args ...interface{}) *sql.Row {
	return db.QueryRow(query, args...)
}

func Exec(query string, args ...interface{}) (sql.Result, error) {
	return db.Exec(query, args...)
}

func Close() {
	if err := db.Close(); err != nil {
		logger.Error("Close: %s", err)
	}
}
