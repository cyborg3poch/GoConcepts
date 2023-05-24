package main

import "fmt"

type Book struct {
	name  string
	price int
}

// Its always good practise to create method based on pointer reciever variable

func (b *Book) PrintBook() {
	fmt.Printf("Book name is : %s \nBook Price : %d", b.name, b.price)
}

//Go automatically sends address of value type struct
func (b *Book) DiscountPrice(discount int) {
	b.price -= discount
}
