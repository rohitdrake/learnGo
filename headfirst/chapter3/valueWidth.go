package main

import "fmt"


func main() {
	fmt.Printf("%12s | %s\n", "Product", "Cost in cents")

	fmt.Println("---------------------------------------------")

	fmt.Printf("%12s | %2d\n", "banana", 9)
	fmt.Printf("%12s | %2d\n", "apple", 12)
	fmt.Printf("%12s | %2d\n", "kiwi", 81)
}