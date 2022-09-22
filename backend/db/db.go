package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go/resst-app/config"
)

const (
	dns string = "root:@Root321@tcp(127.0.0.1:3306)/GO_DB?parseTime=true&charset=utf8"
)

type database struct {
	db *sql.DB
}



var DB = &database{}
var err error

func DBinit(app *config.Application) {
	DB.db, err = sql.Open("mysql", dns)

	if err != nil {
		app.Logger.Println("unable to connect to db", err)

	} else {
		app.Logger.Println("db connected successfully")
	}
}

func (db *database) DBPing() error {

	err := db.db.Ping()

	if err != nil {
		return err
	} else {
		return nil
	}
	
}



