package main

import (
	"fmt"
	"reflect"
)

func main() {
	number := 6
	fmt.Println(&number)
	fmt.Println(reflect.TypeOf(&number))
	fmt.Println(number)

	var myInt int
	fmt.Println(&myInt)
	fmt.Println(reflect.TypeOf(&myInt))
	var myBool bool
	fmt.Println(&myBool)
	fmt.Println(reflect.TypeOf(&myBool))
	var myFloat float64
	fmt.Println(&myFloat)
	fmt.Println(reflect.TypeOf(&myFloat))
}