package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func connectToDatabase(address string, user string, password string, database string, sslMode string) *sql.DB {
	connStr := "postgres://" + user + ":" + password + "@" + address + "/" + database + "?sslmode=" + sslMode
	db, err := sql.Open("postgres", connStr)
	errorOccurred(err, true)
	return db
}
