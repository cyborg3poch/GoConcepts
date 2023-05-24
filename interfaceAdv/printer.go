package main

type printer interface {
	print()
}

// defining list type which is a slice of interface it will be able to store intaces of struct who have also implemented printer interface
type list []printer

// // Implemented Printer Interface for list type as well
func (l list) print() {
	if len(l) == 0 {
		println("Nothin in the store ")
	}

	for _, it := range l {
		it.print()
	}

}
