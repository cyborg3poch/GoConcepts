package main

import (
	"fmt"
	"sync"
)

var msg string

var wg sync.WaitGroup

func updateMessage(str string, m *sync.Mutex) {

	defer wg.Done()
	m.Lock()
	msg = str
	m.Unlock()
}

/*
--->"Output is still not defined"<----
it can either be Hello Universe or Hello Cosmos BUT
WE ARE ACCESSING THE DATA (MSG) SAFELY
THIS IS A THREAD SAFE IMPLEMENTATION OF THE PROBLEMWG PROGRAM

check for race with "go run -race .
*/

func main() {

	msg = "Hello World"

	var mutex sync.Mutex

	wg.Add(2)
	go updateMessage("Hello Universe!", &mutex)
	go updateMessage("Hello Cosmos!", &mutex)
	wg.Wait()

	fmt.Println(msg)

}
