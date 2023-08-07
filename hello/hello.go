package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	firstName := "Radvil"
	lastName := "Laode"
	fullName := fmt.Sprintf("%v %v", firstName, lastName)
	message := greetings.Hello(fullName)
	fmt.Println(message)
}
