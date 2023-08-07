package greetings

import "fmt"

func Hello(name string) string {
  message := fmt.Sprintf("Hello %v, welcome to the jungle!", name)
  return message
}
