package main

import (
	"github.com/streadway/amqp"
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

// createUser forges an SQL query to send to Postgres database, then
// returns if the process was successful with either UUID of a new user or
// and error alongside it.
func createUser(username string, password string, discord string) (bool, string) {

	if len(username) > 18 || len(username) < 4 {
		return false, "Username has incorrect length."
	}

	if len(password) > 64 || len(password) < 6 {
		return false, "Password has incorrect length."
	}

	discordMatch, err := regexp.MatchString(`^.+\#[0-9]{4}$`, discord)
	if errorOccurred(err, false) || !discordMatch {
		return false, "Discord handle is incorrect."
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if errorOccurred(err, false) {
		return false, "Error hashing password"
	}

	discordConfirm := getRandomString(16)
	tx, err := Database.Begin()
	if errorOccurred(err, false) {
		return false, "Database error (Transaction begin)"
	}

	stmt, err := tx.Prepare(`insert into users (name, discord, discordConfirm, password) values ($1, $2, $3, $4) returning id`)
	if errorOccurred(err, false) {
		return false, "Database error (Transaction prepare)"
	}
	defer stmt.Close()

	var userID string
	err = stmt.QueryRow(username, discord, discordConfirm, hashedPassword).Scan(&userID)
	if errorOccurred(err, false) {
		return false, "Database error (Transaction query, scan)"
	}

	err = tx.Commit()
	if errorOccurred(err, false) {
		return false, "Database error (Transaction commit)"
	}

	RabbitMQChannel.Publish("", "SendDiscordValidationMessage", false, false, amqp.Publishing{
		ContentType: "text/plain", Body: []byte(discord + ":" + discordConfirm)})

	return true, userID

}

func confirmDiscord(userID string, discordConfirm string) (bool, string) {

	tx, err := Database.Begin()
	if errorOccurred(err, false) {
		return false, "Database error (Transaction begin)"
	}

	stmt, err := tx.Prepare(`select discordconfirm from users where id = $1`)
	if errorOccurred(err, false) {
		return false, "Database error (Transaction prepare)"
	}

	var discordConfirmDB string
	err = stmt.QueryRow(userID).Scan(&discordConfirmDB)
	if errorOccurred(err, false) {
		return false, "User does not exist"
	}

	if discordConfirmDB != discordConfirm {
		return false, "Confirmation code is not correct"
	}
	stmt.Close()
	stmt, err = tx.Prepare(`update users set discordconfirm = null where id = $1`)
	if errorOccurred(err, false) {
		return false, "Database error (Transaction prepare 2)"
	}
	defer stmt.Close()

	row := stmt.QueryRow(userID)
	if errorOccurred(row.Err(), false) {
		return false, "Database error (Transaction query 2)"
	}

	err = tx.Commit()
	if errorOccurred(err, false) {
		return false, "Database error (Transaction commit)"
	}

	return true, "Discord confirmed"

}
