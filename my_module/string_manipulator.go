package my_module

import (
	"errors"
	"fmt"
)

// Hello it's function name, and start with capital letter because who call this function not in this package.
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf("Hello %v", name)
	return message, nil
}
