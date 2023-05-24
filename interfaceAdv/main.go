package main

func main() {

	junglebook := Book{name: "The JungleBook", price: 6767}
	mineCraft := Game{name: "MicroSoft MineCraft", price: 8989}

	var store list

	store = append(store, &junglebook, mineCraft)
	store.print()

}
