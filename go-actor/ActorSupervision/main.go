package main

import (
	"fmt"

	console "github.com/asynkron/goconsole"
	"github.com/asynkron/protoactor-go/actor"
)

// Define Struct for message
type SimpleMessage struct{ message string }

// Define Struct for actor
type ParentActor struct{}

func (state *ParentActor) Receive(context actor.Context) {
	switch recmsg := context.Message().(type) {
	case SimpleMessage:
		fmt.Printf("Hello %v", recmsg.message)

	}
}

// Define Struct for Child actor
type ChildtActor struct{}

// Defining recieves method on the actor it  a message and prints it if that is of type hello
// Defualt Behavior of the actor is recieve
func (state *ChildtActor) Receive(context actor.Context) {
	// switch based on what type of message we Recieve
	switch recmsg := context.Message().(type) {
	case *actor.Started:
		fmt.Printf(" Hello Actor is started / initialised")
	case *actor.Stopping:
		fmt.Printf(" Hello Actor is about to be stopped")
	case *actor.Stopped:
		fmt.Printf(" Hello Actor is stopped")
	case *actor.Restarting:
		fmt.Printf(" Hello Actor is restarting")
	case SimpleMessage:
		fmt.Printf(" Hello %v", recmsg.message)

	}
}

func NewParentActor() actor.Actor {
	return &ParentActor{}
}

func NewChildActor() actor.Actor {
	return &ChildtActor{}
}

// decider function --- this decides what to do when one or more child fails

func decider(reson interface{}) actor.Directive {
	fmt.Printf("Handling failure of child Reson: %v", reson)
	return actor.RestartDirective
	//return actor.DefaultDecider()
	//return actor.EscalateDirective
	//return actor.ResumeDirective
	//return actor.StopDirective

}

func main() {

	// how often is child actor alowed to fail 10 = max no of retries , 100 = time duration
	supervisor := actor.NewOneForOneStrategy(10, 100, decider)

	//Creating an context which will create an actor system and will be used to send message
	context := actor.NewActorSystem().Root
	props := actor.PropsFromProducer(NewParentActor, actor.WithSupervisor(supervisor))
	pid := context.Spawn(props)

	context.Send(pid, SimpleMessage{message: "Shivank"})

	console.ReadLine()
}
