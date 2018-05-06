package dao

import (
	"fmt"
	"log"
)

const (
	insertMessage = "INSERT INTO messages (username, message) VALUES(?, ?)"
	selectMessages = "SELECT username, message FROM messages ORDER BY timestamp ASC LIMIT 10"
)

func InsertMessage(username string, message string) error {
	stmt, err := db.Prepare(insertMessage)
	if err != nil {
		return fmt.Errorf("error preparing insert user statement [%s]", err)
	}

	defer stmt.Close()

	res, err := stmt.Exec(username, message)
	if err != nil {
		return fmt.Errorf("error preparing insert user statement [%s]", err)
	}

	id, _ := res.LastInsertId()

	log.Print("successfully inserted message " + string(id))
	return nil
}

func GetMessage() ([]string, error) {
	stmt, err := db.Prepare(selectMessages)
	if err != nil {
		return nil, fmt.Errorf("at prepare error: %s", err)
	}

	defer stmt.Close()

	var messages []string
	err = stmt.QueryRow().Scan(&messages)
	if err != nil {
		return nil, fmt.Errorf("at select error: %s", err)
	}

	return messages, nil
}