package main

import "fmt"

func main() {
	var marks [3]int
	marks[0]++
	marks[0]++
	marks[2]++

	fmt.Println(marks[0])
	fmt.Println(marks[1])

	var full_name [2]string

	full_name[1] = "Bhardwaj"

	fmt.Println(full_name[0])
	fmt.Println(full_name[1])

}