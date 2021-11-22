package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var Connection *sql.DB

func init() {
	_db, err := sql.Open("mysql", "root:Princess527#@/gongjiayun")
	if err != nil {
		log.Fatalf("%+v", err)
	}

	Connection = _db
}