package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/jackliu97/chatter/dao"
	"github.com/jackliu97/chatter/data"
)

type user struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewUser(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)

	var u user
	json.Unmarshal(body, &u)

	user, err := data.MakeUser(u.Username, u.Password)
	if err != nil {
		log.Print(err)
		JsonWriter(w, &Response{
			Code:  http.StatusBadRequest,
			Error: fmt.Sprintf("User [%s] failed to create", u.Username),
		})
		return
	}

	err = dao.InsertUser(user)
	if err != nil {
		log.Print(err)
		JsonWriter(w, &Response{
			Code:  http.StatusBadRequest,
			Error: fmt.Sprintf("User [%s] failed to create", u.Username),
		})
		return
	}

	JsonWriter(w, &Response{
		Code: http.StatusCreated,
	})
}

func LogIn(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)

	var u user
	json.Unmarshal(body, &u)

	err := dao.VerifyUser(u.Username, u.Password)
	if err != nil {
		log.Println(err)
		JsonWriter(w, &Response{
			Code:  http.StatusUnauthorized,
			Error: fmt.Sprintf("Invalid user [%s]", u.Username),
		})
		return
	}

	//TODO: On success, generate session id.

	JsonWriter(w, &Response{
		Code: http.StatusOK,
	})
}
