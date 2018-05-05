package main

import (
	"chatter/handlers"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {
	addr := "127.0.0.1:8080"
	r := mux.NewRouter()
	r.HandleFunc("/chat", handlers.GetChatHandler).Methods("GET")
	r.HandleFunc("/chat", handlers.PostChatHandler).Methods("POST")

	srv := &http.Server{
		Handler:      r,
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Starting chat server on " + addr + "...")
	log.Fatal(srv.ListenAndServe())
}
