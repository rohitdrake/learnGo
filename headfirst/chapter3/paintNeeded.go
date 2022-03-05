package main

import "fmt"

func paintNeeded(w float64, h float64) (float64, error) {
	if( w < 0) {
		return 0, fmt.Errorf("A width of %0.2f is invalid", w)
	}

	if( h < 0) {
		return 0, fmt.Errorf("A height of %0.2f is invalid", h)
	}

	area := w*h 
	return area/10.0, nil
	// fmt.Printf("%.2f liters needed\n", area/10.0)
}


func main() {
	amount, err := paintNeeded(4.2, -3.0)
	fmt.Println(err)
	fmt.Println(amount)
	amount, err = paintNeeded(4.2, 5.0)
	fmt.Println(err)
	fmt.Println(amount)
}