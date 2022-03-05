package main

import "fmt"

func multipleReturns() (int, bool, string)	{
	return 12, true, "Alan Turing"
}

func main() {
	number, truth, name := multipleReturns()
	fmt.Println(number, truth, name)
}