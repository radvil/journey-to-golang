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

	// request a greeting message.
	message, err := greetings.Hello("")
	// if an error was returned, print it to the console
	// then exit the program
	if err != nil {
		log.Fatal(err)
	}

	// if no error was returned print the message instead.
	fmt.Println(message)
}
