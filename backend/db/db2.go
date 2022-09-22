package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

const (
	con string = "root:@Root321@tcp(127.0.0.1:3306)/GO_DB?parseTime=true&charset=utf8"
)

func DatabaseInit() (*sql.DB, error) {
	db, err := sql.Open("mysql", con)

	if err != nil {
		return db, err

	}

	return db, nil

}
