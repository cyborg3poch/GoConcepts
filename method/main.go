package main

func main() {
	book := Book{name: "Getting on Nergves", price: 123}
	book.PrintBook()
	book.DiscountPrice(20)
	book.PrintBook()

}
