package main

import "fmt"

func main() {
	array := [3]int{1, 2, 3}
	incArray(&array)
	fmt.Println(array)

}

func incArray(arr *[3]int) {
	for i := range arr {
		(*arr)[i]++ // we dont have to use derefrencing operator (Go does that automatically) though we can use it as well
	}

}
