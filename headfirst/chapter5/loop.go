package main

import "fmt"

func main()	{
	notes := [7]string{"Sa", "Re", "Ga", "Ma", "Pa", "Da", "NiSa"}

	// safest way

	for index, notes := range notes {
		fmt.Println(index, notes)
	}

	for i:=0; i<len(notes); i++	{
		fmt.Println(notes[i])
	}

	// unsafe: will cause panic
	for i:=0; i<=len(notes); i++	{
		fmt.Println(notes[i])
	}


}