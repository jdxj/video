package model

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dsnFormat = "%s:%s@tcp(%s)/%s?loc=Local&parseTime=true"
)

var (
	mysql *sql.DB
)

func InitDB(user, password, host, dbname string) error {
	dsn := fmt.Sprintf(dsnFormat, user, password, host, dbname)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	mysql = db
	return db.Ping()
}
