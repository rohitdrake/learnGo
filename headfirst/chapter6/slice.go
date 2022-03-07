package main

import "fmt"

func main() {
	var notes []string
	notes = make([]string, 7)
	notes[0] = "sa"
	append = append(notes, "re")
	fmt.Println(notes)

	var intSlice []int 
	intSlice = append(intSlice, 27)
	fmt.Println(intSlice)

	var register []string
	register = make([]string, 7)
	register[0] = "rohit"
	register = append(register, "bhardwaj")
	fmt.Println(register)
}