package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
	"github.com/jackliu97/chatter/dao"
	"github.com/jackliu97/chatter/data"
)

var (
	clients   = make(map[*websocket.Conn]bool)
	broadcast = make(chan data.Message)
	upgrader  = websocket.Upgrader{}
)

func init() {
	fmt.Println("listening to broadcasts...")
	go messages()
}

func GetMessages(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	pageNum, err := strconv.Atoi(queryValues.Get("page"))
	if err != nil {
		pageNum = 0
	}

	pageSize, err := strconv.Atoi(queryValues.Get("size"))
	if err != nil {
		pageSize = 10
	}

	results, err := dao.GetMessage(pageNum, pageSize)
	if err != nil {
		JsonWriter(w, &Response{
			Code: http.StatusBadRequest,
		})
		return
	}

	JsonWriter(w, &Response{
		Code: http.StatusOK,
		Data: results,
	})
}

func Connections(w http.ResponseWriter, r *http.Request) {
	// connect a client to ws.
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("error at upgrader " + err.Error())
		return
	}

	defer ws.Close()

	clients[ws] = true

	for {
		var msg data.Message
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

		data.ParseMessage(&msg)

		// save message to db
		go dao.InsertMessage(&msg)

		// broadcast to all clients
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
