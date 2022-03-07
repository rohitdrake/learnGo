package main

import "fmt"


func main() {
	var text [3]string = [3]string{
		"I am first sentence",
		"Yo! This is second",
		"Hello, I am the third one",
	}

	fmt.Println(text[2])
}
