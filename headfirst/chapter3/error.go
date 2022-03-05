package main

import "fmt"

func main() {
	err := fmt.Errorf("A height of %0.2f is invalid", -2.33333)
	fmt.Println(err)
	fmt.Println(err.Error())
}


