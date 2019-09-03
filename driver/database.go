package driver

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var connection *sqlx.DB

func GetConnection() *sqlx.DB {
	connection, err := sqlx.Connect("sqlite3", "ijahstore.sqlite3")
	if err != nil {
		log.Print(err)
		panic("Connection can't be made")
	}

	return connection
}
