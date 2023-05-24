package main

import "fmt"

type Book struct {
	name  string
	price int
}

// Implemented Printer Interface
func (book *Book) print() {
	fmt.Println("Name of book is : " + book.name)

}
