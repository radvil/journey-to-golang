package main

import (
  "fmt"
  "example.com/greetings"
)

func main() {
  var firstName string = "Radvil"
  var lastName string = "Laode"
  fullName := fmt.Sprintf("%v %v", firstName, lastName)
  message := greetings.Hello(fullName)
  fmt.Println(message)
}
