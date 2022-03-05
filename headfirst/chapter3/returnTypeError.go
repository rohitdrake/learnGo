package main

import "fmt"

// won't compile
// invalid return type
func double(number float64) float64 {
	fmt.Println("Inside double")
	return int(number * 2)
}