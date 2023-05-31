package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to upgrade connection to WebSocket:", err)
		return
	}
	defer conn.Close()

	for {
		// Read message from client
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message from WebSocket:", err)
			break
		}

		// Process the received message
		log.Printf("Received message: %s\n", message)

		// Send response back to the client
		err = conn.WriteMessage(websocket.TextMessage, []byte("Server received: "+string(message)))
		if err != nil {
			log.Println("Error sending message to WebSocket:", err)
			break
		}
	}
}

func main() {
	http.HandleFunc("/ws", handleWebSocket)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
