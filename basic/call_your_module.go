package main

import (
	"fmt"
	"log"
	"my_module"
)

func main() {
	log.SetPrefix("String manipulator: ")
	log.SetFlags(0) // Disable printing time, source file and line number in log.

	message, err := my_module.Hello("")
	if err != nil {
		log.Fatal(err) // Print error and stop the program
	}

	fmt.Println(message)
}
