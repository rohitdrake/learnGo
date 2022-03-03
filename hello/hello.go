package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Rohit")
	fmt.Println(message)
}