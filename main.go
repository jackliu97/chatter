package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jackliu97/chatter/handlers"
	"net/http"
	"github.com/jackliu97/chatter/dao"
	"github.com/jackliu97/chatter/config"
	"github.com/spf13/viper"
)

func main() {
	fmt.Println("Init Config...")
	config.InitConfig()

	port := fmt.Sprintf(":%d", viper.Get("port"))
	r := mux.NewRouter()

	r.HandleFunc("/user", handlers.NewUser).Methods("PUT")
	r.HandleFunc("/login", handlers.LogIn).Methods("POST")

	r.Handle("/chat", handlers.AuthenticationMiddleware(handlers.GetChatHandler)).Methods("GET")
	r.Handle("/chat", handlers.AuthenticationMiddleware(handlers.PostChatHandler)).Methods("POST")

	fmt.Println("Opening db connection...")
	dao.Init()

	fmt.Println("Starting chat server on port " + port + "...")
	http.ListenAndServe(port, r)
}
