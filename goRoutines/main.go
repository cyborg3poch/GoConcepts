package main

import (
	"fmt"
	"sync"
)

// We're passing a pointer to wg beacuse its strongly not advised to send a copy of wait group
func print(str string, wg *sync.WaitGroup) {
	fmt.Println(str)
	defer wg.Done()
}

func main() {

	var wg sync.WaitGroup

	list := []string{"one", "two", "three", "four", "five", "six", "seven"}

	wg.Add(len(list))

	for i, _ := range list {

		go print(list[i], &wg)
	}
	wg.Wait()

	wg.Add(1)
	print("I run after wards", &wg)

}
