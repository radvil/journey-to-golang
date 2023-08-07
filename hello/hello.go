package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// set the prop for the predefined logger
	// including log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("GREETINGS: ")
	log.SetFlags(0)

	// A slice of names.
	names := []string{"Traxex", "Injoker", "Nevermore"}

	// request greeting message for the name
	messages, err := greetings.Hellos(names)
	// if an error was returned, print it to the console
	// then exit the program
	if err != nil {
		log.Fatal(err)
	}

	// if no error was returned print the returned map of
	// messages to the console.
	fmt.Println(messages)
}
