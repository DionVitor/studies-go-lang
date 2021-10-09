package main

import (
	"fmt"
)

func main() {
	var array [2]string
	array[0] = "Hello"
	array[1] = "World"
	fmt.Println(array)
	fmt.Println("-----------------------")

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	slice := primes[1:4]
	fmt.Println(slice)
	fmt.Println("-----------------------")

	// Note that slice are like references to arrays
	primes[1] = 0
	fmt.Println(primes)
	fmt.Println(slice)
	fmt.Println("-----------------------")

	// Slice literals
	q := []int{2, 4}
	fmt.Println(q)

	r := []bool{true, false, true}
	fmt.Println(r)

	s := []struct{
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{4, false},
	}
	fmt.Println(s)
	fmt.Println("-----------------------")

	// Slice length and capacity -  you can extend a slice and re-slicing it
	arr := []int{2, 3, 4, 5, 6, 7}
	printSlice(arr)

	arr = arr[:0]
	printSlice(arr)

	arr = arr[:4]
	printSlice(arr)

	arr = arr[2:]
	printSlice(arr)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}