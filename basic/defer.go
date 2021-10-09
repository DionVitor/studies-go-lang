package main

import "fmt"

func main() {
	//defer fmt.Println("World!")
	//
	//fmt.Printf("Hello ")
	stackingDefers()
}

func stackingDefers() {
	fmt.Println("counting")

	// Deferred functions calls are pushed onto a stack.
	// When a function returns, it deferred calls are executed in last-in-first-out order.
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
