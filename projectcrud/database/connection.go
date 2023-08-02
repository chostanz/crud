package database

import (
	"fmt"

	sqlx "github.com/jmoiron/sqlx"
)

var DB = Connection()

func Connection() *sqlx.DB {
	db, err := sqlx.Connect("postgres", "user=postgres password=12345 dbname= db_users sslmode=disable")

	if err = db.Ping(); err != nil {
		fmt.Println(err)
	}
	return db
}
