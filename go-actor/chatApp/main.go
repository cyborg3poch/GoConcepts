package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/asynkron/protoactor-go/actor"
)

// ---- Messages ----

// StartChatRoom message instructs the server to start a new chat room actor.
type StartChatRoom struct {
	RoomID string
}

// JoinRoom message is sent by clients to join a chat room.
type JoinRoom struct {
	Sender  *actor.PID
	Message string
}

// LeaveRoom message is sent by clients to leave a chat room.
type LeaveRoom struct {
	Sender *actor.PID
}

// ChatMessage message represents a message sent by a client in a chat room.
type ChatMessage struct {
	Sender  *actor.PID
	Message string
}

// ChatRoomActor represents an actor responsible for managing a chat room.
type ChatRoomActor struct {
	RoomID  string
	clients map[*actor.PID]bool
}

// Receive handles incoming messages for the ChatRoomActor.
func (state *ChatRoomActor) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *JoinRoom:
		state.clients[msg.Sender] = true
		state.Broadcast(fmt.Sprintf("User joined: %s", msg.Sender.String()))

	case *LeaveRoom:
		delete(state.clients, msg.Sender)
		state.Broadcast(fmt.Sprintf("User left: %s", msg.Sender.String()))

	case *ChatMessage:
		state.Broadcast(fmt.Sprintf("%s: %s", msg.Sender.String(), msg.Message))

	case *actor.Stopping:
		// Cleanup when the actor is stopped
		log.Printf("Chat room %s is stopped.", state.RoomID)

	}
}

// Broadcast sends a message to all clients in the chat room.
func (state *ChatRoomActor) Broadcast(message string) {
	for client := range state.clients {
		client.Tell(&ChatMessage{Sender: nil, Message: message})
	}
}

// ---- Main ----

func main() {
	// Create the root actor system
	system := actor.NewActorSystem()

	// Create a root supervisor actor
	rootProps := actor.PropsFromProducer(func() actor.Actor {
		return &ChatRoomSupervisorActor{}
	}).WithMiddleware(
		middleware.Logger,
		middleware.Recover,
	)
	rootContext := system.Root
	rootContext.Spawn(rootProps)

	// Start the chat room
	rootContext.Send(rootContext.Root, &StartChatRoom{RoomID: "general"})

	// Wait for termination signal (Ctrl+C)
	waitForTerminationSignal()
}

// ---- Supervisor Actor ----

// ChatRoomSupervisorActor represents the supervisor actor responsible for managing chat room actors.
type ChatRoomSupervisorActor struct {
}

// Receive handles incoming messages for the ChatRoomSupervisorActor.
func (state *ChatRoomSupervisorActor) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *StartChatRoom:
		roomProps := actor.PropsFromProducer(func() actor.Actor {
			return &ChatRoomActor{
				RoomID:  msg.RoomID,
				clients: make(map[*actor.PID]bool),
			}
		})
		ctx.Spawn(roomProps)
		log.Printf("Chat room %s started.", msg.RoomID)
	}
}

// ---- Utility ----

// waitForTerminationSignal waits for a termination signal (Ctrl+C) to gracefully shutdown the application.
func waitForTerminationSignal() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	fmt.Println("\nTermination signal received. Shutting down...")
	time.Sleep(1 * time.Second)
	os.Exit(0)
}
