package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("glady")
	fmt.Println(message)
}
