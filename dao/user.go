package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jackliu97/chatter/pogo"
	"log"
	"fmt"
)

const (
	insertUser = "INSERT INTO users (username, password) VALUES(?, ?)"
	selectUser = "SELECT username, password FROM users WHERE username = ?"
)

func InsertUser(user *pogo.User) {
	stmt, err := db.Prepare(insertUser)
	if err != nil {
		log.Fatal("error preparing insert user statement " + err.Error())
	}

	defer stmt.Close()
	res, err := stmt.Exec(user.GetUsername(), user.GetPassword())
	if err != nil {
		log.Fatal("error preparing insert user statement " + err.Error())
	}

	id, _ := res.LastInsertId()

	log.Print("successfully inserted user " + string(id))
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

	user, err := pogo.MakeUser(usernameResult, passwordResult)
	if err != nil {
		return fmt.Errorf("at MakeUser username: [%s] error: %s", username, err)
	}

	return user.VerifyPassword(password)
}