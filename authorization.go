package main

import "time"

// generateAPIKey forges an SQL query to send to Postgres database, then
// returns if the process was successful with either generated key or
// and error alongside it.
func generateAPIKey(userID string, expiryDateUnix int64) (bool, string) {

	tx, err := Database.Begin()
	if errorOccurred(err, false) {
		return false, "Database error (Transaction begin)"
	}

	stmt, err := tx.Prepare(`select name from users where id = $1`)
	if errorOccurred(err, false) {
		return false, "Database error (Transaction prepare 1)"
	}
	var userName string
	err = stmt.QueryRow(userID).Scan(&userName)
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
