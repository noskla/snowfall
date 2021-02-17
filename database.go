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

func createTablesIfNotExists(database *sql.DB) {
	err := database.Ping()
	errorOccurred(err, true)

	queries := []string{
		"create extension if not exists \"uuid-ossp\"",
		"create table if not exists rooms (id uuid default uuid_generate_v4(), name varchar not null, description text, primary key (id))",
		"create table if not exists events (id uuid default uuid_generate_v4(), roomID uuid, startDate date not null, endDate date not null, organizerID uuid[], description text, primary key (id))",
		"create table if not exists stands (id uuid default uuid_generate_v4(), name varchar not null, description text, offers uuid[], owners uuid[])",
		"create table if not exists standOffers (id uuid default uuid_generate_v4(), name varchar not null, price int default 50)",
		"create table if not exists users (id uuid default uuid_generate_v4(), name varchar not null, discord varchar not null, discordConfirm varchar, password varchar not null, admin boolean default false)",
		"create table if not exists keys (authKey varchar(26) unique not null default substring(md5(random()::text), 0, 25), owner uuid not null, expiresAfter date default null)",
	}

	for _, query := range queries {
		_, err := database.Exec(query)
		errorOccurred(err, true)
	}

	log.Println("Tables initialized.")
}
