package main

import "fmt"



func main() {
	var notes [7]string

	notes[0]="Sa"
	notes[1]="Re"
	notes[2]="Ga"

	fmt.Println(notes[0], notes[1], notes[2])

	var prime [5]int
	prime[0]=3
	prime[1]=5
	prime[2]=7

	fmt.Println(prime[0], prime[1], prime[2])
}