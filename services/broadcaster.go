package services

import (
	"sync"

	"github.com/gorilla/websocket"
)

type SafeWebSocket struct {
	Conn  *websocket.Conn
	mutex sync.Mutex
}

var Clients map[*SafeWebSocket]bool

var ClientConnection = make(map[*SafeWebSocket]bool)

func (ws *SafeWebSocket) WriteMessage(messageType int, data []byte) error {
	ws.mutex.Lock()
	defer ws.mutex.Unlock()
	return ws.Conn.WriteMessage(messageType, data)
}

func SendToAllClients(message []byte) {
	for conn := range ClientConnection {
		conn.WriteMessage(websocket.TextMessage, message)
	}
}
