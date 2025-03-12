package adapters

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool) // Lista de clientes conectados
var broadcast = make(chan string) // Canal para enviar mensajes a los clientes

// WebSocketAdapter maneja las conexiones WebSocket
type WebSocketAdapter struct {
	Upgrader websocket.Upgrader
}

func NewWebSocketAdapter() *WebSocketAdapter {
	return &WebSocketAdapter{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true // Permitir todas las conexiones
			},
		},
	}
}

// Maneja la conexión WebSocket
func (adapter *WebSocketAdapter) HandleWebSocketConnection(conn *websocket.Conn) {
	defer conn.Close()
	clients[conn] = true

	for {
		_, _, err := conn.ReadMessage() // Mantener la conexión abierta
		if err != nil {
			delete(clients, conn) // Eliminar cliente cuando se cierra
			break
		}
	}
}

// Enviar un mensaje a todos los clientes WebSocket conectados
func (adapter *WebSocketAdapter) BroadcastMessage(message string) {
	for client := range clients {
		err := client.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			log.Println(err)
			client.Close()
			delete(clients, client)
		}
	}
}

// Inicia la escucha de mensajes y los transmite a los clientes
func (adapter *WebSocketAdapter) Start() {
	for message := range broadcast {
		adapter.BroadcastMessage(message)
	}
}
