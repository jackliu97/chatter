package handlers

import "net/http"

var GetChatHandler = http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
	message := "Getting existing chat"
	w.Write([]byte(message))
})

var PostChatHandler = http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
	message := "Posting new chat"
	w.Write([]byte(message))
})
