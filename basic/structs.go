package main

import "fmt"

// Structs is a collection of fields.

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	fmt.Println(v.X)
	fmt.Println(v.Y)
}
