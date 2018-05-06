package handlers

import (
	"net/http"
	"github.com/jackliu97/chatter/dao"
	"github.com/jackliu97/chatter/pogo"
	"log"
	"fmt"
)

type User struct {
	Username string
	Password string
}

func NewUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	username := r.FormValue("username")
	user, err := pogo.MakeUser(username, r.FormValue("password"))
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("User [%s] failed to create", username)))
		return
	}

	dao.InsertUser(user)

	w.Write([]byte(fmt.Sprintf("User [%s] successfully created", username)))
}

func LogIn(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	username := r.FormValue("username")
	err := dao.VerifyUser(username, r.FormValue("password"))

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(fmt.Sprintf("Invalid user [%s]", username)))
		return
	}

	w.Write([]byte(fmt.Sprintf("User [%s] logged in successfully!", username)))
}