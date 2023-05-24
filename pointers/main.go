package main

import "fmt"

func main() {

	// When * is put before a type it becomes a type declaration of pointer of that type

	//var p *int
	z := 1
	p := &z // Auto declare and initialization of int pointer

	fmt.Printf("%d\n", *p) // when * is put before a pointer variable it becomes a derefrencing operator

	*p = 77
	fmt.Printf("%d", z)

}
