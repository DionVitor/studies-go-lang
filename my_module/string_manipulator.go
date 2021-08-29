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


// Hellos returns a map that associates each of the named people with a hello message.
func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}
