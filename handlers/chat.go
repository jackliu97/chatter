package handlers

import (
	"github.com/gorilla/websocket"
	"github.com/jackliu97/chatter/pogo"
	"log"
	"net/http"
)

var (
	clients = make(map[*websocket.Conn]bool)
	broadcast = make(chan pogo.Message)
	upgrader = websocket.Upgrader{}
)

func InitChat() {
	go messages()
}

func Connections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("error at upgrader " + err.Error())
		return
	}

	defer ws.Close()

	clients[ws] = true

	for {
		var msg pogo.Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error at reading : %s", err)
			delete(clients, ws)
			break
		}

		broadcast <- msg
	}
}

func messages() {
	for {
		msg := <-broadcast
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error at handling message %s", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
