package main

import (
	"fmt"
	"sync"
)

var msg string

var wg sync.WaitGroup

func updateMessage(str string, wg *sync.WaitGroup) {
	defer wg.Done()
	msg = str
}

/*
Running this program only either of the mesasge gets printed (in no defined order) because there is a race condition
Both of the go routine try(race) to update the msg variable
Also since we're using wait groups the order of the execution is not gauranteed

Refer to mutex.go in same package for a resolve of this probem using mutex from sync.Mutex
*/

func main() {

	msg = "Hello World"

	wg.Add(2)

	go updateMessage("Hello Universe!", &wg)
	go updateMessage("Hello Cosmos!", &wg)
	wg.Wait()

	fmt.Println(msg)

}
