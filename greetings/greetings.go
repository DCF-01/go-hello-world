package greetings

import (
	"fmt"
	"errors"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Error: argument {name} cannot be empty.")
	}

	message := fmt.Sprintf("Hi %v, welcome.", name)
	return message, nil
}