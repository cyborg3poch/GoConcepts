package main

import "fmt"

type greet struct {
	Name string
	Msg  string
}

func (g *greet) greeter() {

	g.Name = "Mayank"
	fmt.Println(g.Name, g.Msg)

}

func main() {

	g := greet{"Shivank", "Kaise ho!"}
	g.greeter()
	fmt.Println(g.Name)

}
