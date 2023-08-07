package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("EMPTY_NAME")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// Hellos returns a map of associates each of the named people
// with a greeting message
func Hellos(names []string) (map[string]string, error) {
	// A map to associate names and messages
	messages := make(map[string]string)
	// Loop through the received slice of names, calling
	// the Hello function to get a message to each name.
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// In the map, associate the retrieved message with the name
		messages[name] = message
	}

	return messages, nil
}

// helper function
// randomFormat returns one of a set of greeting message,
// the returned message is selected on random.
func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hi, %v Welcome!",
		"Nice to see you, %v",
		"Hi %v, Well met!",
	}

	// return on of the random formatted message
	return formats[rand.Intn(len((formats)))]
}
