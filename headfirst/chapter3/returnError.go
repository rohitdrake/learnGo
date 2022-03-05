package main

import "fmt"

// func must end with return statement
func double(number float64) float64 {
	product := number*2
	return product
	fmt.Println(product)
}

func main() {
	double(12)
}