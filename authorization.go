package main

import "time"

// GenerateAPIKey forges an SQL query to send to Postgres database, then
// returns if the process was successful with either generated key or
// and error alongside it.
func GenerateAPIKey(userID string, expiryDateUnix int64) (bool, string) {

	tx, err := Database.Begin()
	if errorOccurred(err, false) {
		return false, "Database error (Transaction begin)"
	}

	stmt, err := tx.Prepare(`select name from users where id = $1`)
	if errorOccurred(err, false) {
		return false, "Database error (Transaction prepare 1)"
	}
	var userName string
	row := stmt.QueryRow(userID)
	if row == nil {
		return false, "User does not exist."
	}

	err = row.Scan(&userName)
	if errorOccurred(err, false) {
		return false, "Database error (Transaction query, scan 1)"
	}
	stmt.Close()

	stmt, err = tx.Prepare(`insert into keys (owner, expiresafter) values ($1, $2) returning authkey`)
	if errorOccurred(err, false) {
		return false, "Database error (Transaction prepare 2)"
	}

	var authKey string
	if expiryDateUnix <= int64(time.Now().Unix()) {
		err = stmt.QueryRow(userID, nil).Scan(&authKey)
	} else {
		err = stmt.QueryRow(userID, expiryDateUnix).Scan(&authKey)
	}
	if errorOccurred(err, false) {
		return false, "Database error (Transaction query, scan 2)"
	}
	defer stmt.Close()

	err = tx.Commit()
	if errorOccurred(err, false) {
		return false, "Database error (Transaction commit)"
	}

	return true, authKey

}

// ValidateAPIKey checks the given key and returns if the query succeeded and
// if true; the permission level assigned to related user.
// 0 - key does not exist, 1 - admin: false, 2 - admin: true
func ValidateAPIKey(apiKey string) (bool, uint8) {

	tx, err := Database.Begin()
	if errorOccurred(err, false) {
		return false, 0
	}

	stmt, err := tx.Prepare(`select users.admin from keys left join users on users.id = keys.owner where keys.authkey = $1`)
	if errorOccurred(err, false) {
		return false, 0
	}
	defer stmt.Close()

	var isAdmin bool
	err = stmt.QueryRow(apiKey).Scan(&isAdmin)
	if errorOccurred(err, false) {
		return false, 0
	}

	permLevel := uint8(1)
	if isAdmin {
		permLevel++
	}

	return true, permLevel

}
