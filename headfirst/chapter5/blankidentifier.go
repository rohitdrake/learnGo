package main

import "fmt"

func main() {
	marks := [5]int{88, 92, 95, 93, 94}

	// not using index
	for _,mark := range marks {
		fmt.Println(mark)
	}

	fmt.Println()

	// not using value
	for subject, _ := range marks {
		fmt.Println(subject)
	}	
}