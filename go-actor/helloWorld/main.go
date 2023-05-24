package main

import (
	"fmt"

	"github.com/asynkron/protoactor-go/actor"
)

// Define Struct for message
type SimpleMessage struct{ message string }

//type OtherMessage struct{ message string }

// Define Struct for actor
type Actor struct{}

// Defining recieves method on the actor it  a message and prints it if that is of type hello
// Defualt Behavior of the actor is recieve-- important
func (state *Actor) Receive(context actor.Context) {

	// switch based on what type of message we Recieve
	switch recmsg := context.Message().(type) {
	case SimpleMessage:
		fmt.Printf(" Hello %v", recmsg.message)
		// case OtherMessage:
		// 	actr.onOtherEvent(context)
	default:
		fmt.Println("i ran ")

	}
}

// Extending the default behavior of the actor

// func (actr *Actor) onOtherEvent(context actor.Context) {
// 	switch recmsg := context.Message().(type) {
// 	case OtherMessage:
// 		fmt.Printf(" Hello %v , The actor is now in a different state", recmsg.message)

// 	}

// }

func NewActor() actor.Actor {
	return &Actor{}
}

func main() {
	//Creating an context which will create an actor system and will be used to send message

	// root context
	context := actor.NewActorSystem().Root
	props := actor.PropsFromProducer(NewActor)

	pid := context.Spawn(props)

	context.Send(pid, SimpleMessage{message: "Shivank"})

	//console.ReadLine()
}
