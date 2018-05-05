package handlers

import "net/http"

func GetChatHandler(w http.ResponseWriter, r *http.Request) {
	message := "Getting existing chat"
	w.Write([]byte(message))
}

func PostChatHandler(w http.ResponseWriter, r *http.Request) {
	message := "Posting new chat"
	w.Write([]byte(message))
}
