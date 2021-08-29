package main

import (
	"fmt"
	"my_module"
)

func main() {
	message := my_module.Hello("Dion")
	fmt.Println(message)
}
