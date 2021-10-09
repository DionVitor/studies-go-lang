package main

import "fmt"

func main() {
	// Make function is to create dynamically-sized arrays
	// Params: array, length and capacity

	a := make([]int, 5)
	pprintSlice("a", a)

	b := make([]int, 0, 5)
	pprintSlice("b", b)
}

func pprintSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
