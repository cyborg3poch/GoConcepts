WebSockets are a communication protocol that provides full-duplex communication channels over a single TCP connection. They enable real-time communication between a client and a server, allowing both parties to send and receive data simultaneously. WebSockets are widely used in applications that require continuous data exchange, such as chat applications, real-time dashboards, and collaborative tools.

In Go, you can use the "github.com/gorilla/websocket" package to work with WebSockets. Here's an example that demonstrates the use of WebSockets in Go:

```go
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
```

In this example, we create an HTTP server that handles WebSocket connections. The `/ws` endpoint is used for WebSocket communication. When a client requests an upgrade to WebSocket, the `handleWebSocket` function is called.

Inside the `handleWebSocket` function, we first upgrade the HTTP connection to a WebSocket connection using the `Upgrader` provided by the `gorilla/websocket` package. We then enter a loop where we continuously read messages from the client using `conn.ReadMessage()`. Once we receive a message, we log it and send a response back to the client using `conn.WriteMessage()`.

To run this example, you need to have the `gorilla/websocket` package installed. You can install it using the following command:

```
go get -u github.com/gorilla/websocket
```

Then, save the example code in a file named `main.go` and run it with the following command:

```
go run main.go
```

Now, you have a WebSocket server running on `localhost:8080/ws` that echoes back any messages it receives from the clients. You can use WebSocket client libraries or browser APIs to connect to this server and exchange real-time data.