package dao

import (
	_ "github.com/go-sql-driver/mysql"

	"fmt"
	"log"

	"github.com/jackliu97/chatter/data"
)

const (
	insertUser = "INSERT INTO users (username, password) VALUES(?, ?)"
	selectUser = "SELECT username, password FROM users WHERE username = ?"
)

func InsertUser(user *data.User) error {
	stmt, err := db.Prepare(insertUser)
	if err != nil {
		return fmt.Errorf("error preparing insert user statement [%s]", err)
	}

	defer stmt.Close()
	res, err := stmt.Exec(user.GetUsername(), user.GetPassword())
	if err != nil {
		return fmt.Errorf("error preparing insert user statement [%s]", err)
	}

	id, _ := res.LastInsertId()

	log.Print("successfully inserted user " + string(id))
	return nil
}

func VerifyUser(username string, password string) error {
	stmt, err := db.Prepare(selectUser)
	if err != nil {
		return fmt.Errorf("at Prepare username: [%s] error: %s", username, err)
	}

	defer stmt.Close()

	var usernameResult string
	var passwordResult string

	err = stmt.QueryRow(username).Scan(&usernameResult, &passwordResult)
	if err != nil {
		return fmt.Errorf("at Scan username: [%s] error: %s", username, err)
	}

	user, err := data.MakeUser(usernameResult, "")
	if err != nil {
		return fmt.Errorf("at MakeUser username: [%s] error: %s", username, err)
	}

	user.SetHashedPassword(passwordResult)

	return user.VerifyPassword(password)
}
