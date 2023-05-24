package main

import "fmt"

type Game struct {
	name  string
	price int
}

// Implemented Printer Interface
func (game Game) print() {
	fmt.Println("Name of game is : " + game.name)

}
