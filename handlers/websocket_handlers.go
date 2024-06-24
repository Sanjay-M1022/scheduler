package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/scheduler/services"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func JobUpdateBroadcaster(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(w, r, nil)
	newClientConn := services.SafeWebSocket{Conn: conn}
	services.ClientConnection[&newClientConn] = true

	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()
	defer delete(services.Clients, &newClientConn)

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		clientEvent := string(message)
		if clientEvent == "trigger_update" {
			go services.SendToAllClients([]byte("update_jobs_list"))
		}
	}
}
