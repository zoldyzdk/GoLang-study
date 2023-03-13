package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
  //Get a freeting message and print it
  message := greetings.Hello("Hewerton")
  fmt.Println(message)
}
