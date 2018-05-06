package main

import (
	_ "github.com/jackliu97/chatter/config"
	_ "github.com/jackliu97/chatter/dao"
	_ "github.com/jackliu97/chatter/handlers"

	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackliu97/chatter/handlers"
	"github.com/spf13/viper"
)

func main() {
	r := mux.NewRouter()

	// static contents
	r.Handle("/", http.FileServer(http.Dir("./public")))
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public/"))))

	// api routing
	r.HandleFunc("/ws", handlers.Connections)
	r.HandleFunc("/user", handlers.NewUser).Methods("POST")
	r.HandleFunc("/login", handlers.LogIn).Methods("POST")
	r.HandleFunc("/messages", handlers.GetMessages).Methods("GET")

	port := fmt.Sprintf(":%d", viper.Get("port"))
	fmt.Println("Starting chat server on port " + port + "...")
	http.ListenAndServe(port, r)
}
