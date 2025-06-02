package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectToDb() *sql.DB {
	connStr := "user=admin dbname=alura_store password=password host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err.Error())
	}
	return db
}
