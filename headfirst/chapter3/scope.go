package main

import "fmt"

var metersPerLitre float64

func paintNeeded(w float64, h float64) {
	area := w*h 
	fmt.Printf("%.2f liters needed\n", area/metersPerLitre)
}


func main() {
	metersPerLitre=10.0
	paintNeeded(4.2, 3.0)
	paintNeeded(5.2, 3.5)
}