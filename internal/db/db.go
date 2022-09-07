package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"time"
)

var (
	db *sql.DB
)

func Init() error {
	var err error
	db, err = sql.Open("mysql", os.Getenv("MYSQL_URL"))
	if err != nil {
		return err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return nil
}

func Disconnect() {
	_ = db.Close()
}
