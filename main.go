package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jackliu97/chatter/config"
	"github.com/jackliu97/chatter/dao"
	"github.com/jackliu97/chatter/handlers"
	"github.com/spf13/viper"
	"net/http"
)

func main() {
	fmt.Println("Init Config...")
	config.InitConfig()

	port := fmt.Sprintf(":%d", viper.Get("port"))
	r := mux.NewRouter()

	// static contents
	r.Handle("/", http.FileServer(http.Dir("./public")))
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public/"))))

	// api routing
	r.HandleFunc("/ws", handlers.Connections)
	r.HandleFunc("/user", handlers.NewUser).Methods("POST")
	r.HandleFunc("/login", handlers.LogIn).Methods("POST")

	dao.Init()
	dao.Seed()
	handlers.InitChat()

	fmt.Println("Starting chat server on port " + port + "...")
	http.ListenAndServe(port, r)
}
