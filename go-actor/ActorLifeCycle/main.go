package main

import (
	"fmt"
	"log"

	console "github.com/asynkron/goconsole"
	"github.com/asynkron/protoactor-go/actor"
)

// Define Struct for actor message
type SimpleMessage struct{ message string }

type OtherMessage struct{ message string }

// Define Struct for actor
type Actor struct{}

/*
Defining recieves method on the actor it  a message and prints it if that is of type hello
1. Defualt Behavior of the actor is recieve
2.Actor lifecycle events are called no matter what
*/
func (actr *Actor) Receive(context actor.Context) {
	// switch based on what type of message we Recieve
	switch recmsg := context.Message().(type) {
	case *actor.Started:
		log.Printf("%v", context)
		fmt.Printf(" Hello Actor is started / initialised")
	case *actor.Stopping:
		fmt.Printf(" Hello Actor is about to be stopped")
	case *actor.Stopped:
		fmt.Printf(" Hello Actor is stopped")
	case *actor.Restarting:
		fmt.Printf(" Hello Actor is restarting")
	case SimpleMessage:
		fmt.Printf(" Hello %v", recmsg.message)
	case OtherMessage:
		actr.onOtherEvent(context)

	}
}

// Extending the default behavior of the actor

func (actr *Actor) onOtherEvent(context actor.Context) {
	switch recmsg := context.Message().(type) {
	case OtherMessage:
		fmt.Printf(" Hello %v , The actor is now in a different state", recmsg.message)

	}

}

func NewActor() actor.Actor {
	return &Actor{}
}

func main() {
	//Creating an context which will create an actor system and will be used to send message
	// context ~ system
	context := actor.NewActorSystem().Root

	// props define the properties
	props := actor.PropsFromProducer(NewActor)
	pid := context.Spawn(props)

	log.Printf("%v", pid)

	//context.Send(pid, OtherMessage{message: "Shivank"})

	//time.Sleep(1 * time.Second)
	//context.Stop(pid)

	console.ReadLine()
}
