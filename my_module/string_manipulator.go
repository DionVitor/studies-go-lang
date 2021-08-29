package my_module

import "fmt"

// Hello it's function name, and start with capital letter because who call this function not in this package.
func Hello(name string) string {
	message := fmt.Sprintf("Hello %v", name)
	return message
}
