package main

import (
	"fmt"
	"strings"
)

func shout(ping chan string, pong chan string) {

	// Putting everything in for loop so that it keep running and ping channel keep listening
	for {
		// getting value from ping channel in s
		s  := <-ping
		// sending reponse in pong channel
		pong <- fmt.Sprintf("TU %s !!", s)

	}

}

func main() {

	ping := make(chan string)
	pong := make(chan string)

	// calling shout function it will run as a seperate go routine and keep listening for message though ping channel

	go shout(ping, pong)

	fmt.Println("Enter some gaali and I will give it back to you | Q to quit")
	for {
		fmt.Println("-> ")
		var userInput string
		_, _ = fmt.Scanln(&userInput)

		if userInput == strings.ToLower("q") {
			break

		}

		// Sending userInput to ping channel which is also of string type
		ping <- userInput

		// initialising response and getting value from pong channel
		resposne := <-pong

		//simply printing out the response now
		fmt.Println(resposne)

	}

	// closeing all channels
	fmt.Println("Closing Channels")
	close(ping)
	close(pong)

}
