package greetings

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("EMPTY_NAME")
	}
	message := fmt.Sprintf("Hello %v, welcome to the jungle!", name)
	return message, nil
}
