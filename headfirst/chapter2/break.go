package main

import "fmt"

func main()	{

	for x:=1; x<=3; x++ {
		fmt.Println("before continue")
		continue 
		fmt.Println("after contiue")
	}

	for y:=1; y<=3; y++ {
		fmt.Println("before break")
		break 
		fmt.Println("after break")
	}

	fmt.Println("outside loop")

}