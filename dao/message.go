package dao

import (
	"fmt"

	"github.com/jackliu97/chatter/data"
	"log"
)

const (
	insertMessage  = "INSERT INTO messages (username, message) VALUES(?, ?)"
	selectMessages = "SELECT username, message FROM (" +
		"SELECT id, username, message FROM messages ORDER BY messages.id DESC LIMIT ?, ?" +
		") as x ORDER BY x.id"
)

func InsertMessage(msg *data.Message) error {
	stmt, err := db.Prepare(insertMessage)
	if err != nil {
		return fmt.Errorf("error preparing insert user statement [%s]", err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(msg.Username, msg.Message)
	if err != nil {
		return fmt.Errorf("error preparing insert user statement [%s]", err)
	}

	return nil
}

// GetMessage gets all message with page / limit. (page starts from 1)
func GetMessage(pageNum int, pageSize int) ([]data.Message, error) {
	stmt, err := db.Prepare(selectMessages)

	var result []data.Message

	// page starts at 1. if 0 or lower is passed, we assume 1.
	if pageNum < 1 {
		pageNum = 1;
	}

	offset := pageSize * (pageNum-1)

	if err != nil {
		return nil, fmt.Errorf("at prepare error: %s", err)
	}

	defer stmt.Close()

	rows, err := stmt.Query(offset, pageSize)
	if err != nil {
		return nil, fmt.Errorf("at select rows error: %s", err)
	}

	defer rows.Close()

	for rows.Next() {
		var m data.Message
		err := rows.Scan(&m.Username, &m.Message)
		if err != nil {
			log.Printf("failed to scan row error: %s", err)
		}

		result = append(result, m)
	}

	return result, nil
}
