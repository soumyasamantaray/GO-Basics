package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointer")

	//Defining variables
	var pointervar int = 06

	fmt.Println("this is variable pointer module", pointervar)

	myNumber := 22
	var ptr = &myNumber

	fmt.Println("Value of actual pointer is ", ptr)
	fmt.Println("Value of actual pointer is ", *ptr)

	*ptr = *ptr + 2
	fmt.Println("New Value is: ", myNumber)

}
