package main

import "fmt"


func status(marks float64) string {
	if(marks < 60) {
		return "Fail"
	}

	return "Pass"
}

func main() {
	fmt.Println(status(61))
	fmt.Println(status(59))
}