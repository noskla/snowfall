package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func connectToDatabase(address string, user string, password string, database string, sslMode string) *sql.DB {
	connStr := "postgres://" + user + ":" + password + "@" + address + "/" + database + "?sslmode=" + sslMode
	db, err := sql.Open("postgres", connStr)
	errorOccurred(err, true)
	log.Println("Connected to PostgreSQL server")
	return db
}
