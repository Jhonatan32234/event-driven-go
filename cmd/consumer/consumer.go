package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/streadway/amqp"
)

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan string)
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

// Función para manejar WebSocket
func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error al establecer conexión WebSocket:", err)
		return
	}
	defer ws.Close()

	clients[ws] = true

	for msg := range broadcast {
		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, []byte(msg))
			if err != nil {
				fmt.Println("Error al enviar mensaje por WebSocket:", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

// Función para consumir mensajes de RabbitMQ desde dos colas
func consumeMessages() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println("Error al conectar con RabbitMQ:", err)
		return
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println("Error al abrir canal:", err)
		return
	}
	defer ch.Close()

	// Escuchar mensajes de `notification_queue`
	msgsNotification, err := ch.Consume(
		"notification_queue", // Cola
		"",                   // Consumer (vacío para generar uno aleatorio)
		true,                 // Auto-Acknowledge
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println("Error al consumir `notification_queue`:", err)
		return
	}

	// Escuchar mensajes de `alert_queue`
	msgsAlert, err := ch.Consume(
		"alert_queue", // Cola
		"",            // Consumer
		true,          // Auto-Acknowledge
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println("Error al consumir `alert_queue`:", err)
		return
	}

	// Goroutine para manejar `notification_queue`
	go func() {
		for msg := range msgsNotification {
			message := string(msg.Body) // Convertir los bytes a string
			fmt.Println("Mensaje desde notification_queue:", message) // Imprimir el mensaje como string JSON
			broadcast <- message
		}
	}()
	
	go func() {
		for msg := range msgsAlert {
			message := string(msg.Body) // Convertir los bytes a string
			fmt.Println("Mensaje desde alert_queue:", message) // Imprimir el mensaje como string JSON
			broadcast <- message
		}
	}()
	

	fmt.Println("[✅] Escuchando mensajes de RabbitMQ en `notification_queue` y `alert_queue`...")
	select {} // Mantiene el proceso activo
}

func main() {
	http.HandleFunc("/ws", handleConnections)
	go consumeMessages()

	fmt.Println("[✅] WebSocket activo en el puerto :5000")
	http.ListenAndServe(":5000", nil)
}
