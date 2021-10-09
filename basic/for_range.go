package main

import "fmt"

func main() {
	pow := []int{1, 2, 4, 8, 16, 32, 62, 128}

	for index, value := range pow {
		fmt.Printf("2^%d = %d\n", index, value)
	}
}
