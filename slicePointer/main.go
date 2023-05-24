package main

import "fmt"

func main() {
	// Following the same technique as array doesn't work for Slice because Slices and Maps are	 referance type
	slice := []int{1, 2, 3}
	incSlice(slice)
	fmt.Println(slice)

}

func incSlice(slice []int) {
	for i := range slice {
		slice[i]++
	}

}
