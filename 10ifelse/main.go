package main

import "fmt"

func main() {
	fmt.Println("Welcome to if else")

	// Adding if else statement
	loginCount := 11
	var result string

	if loginCount < 10 {
		result = "regular user"
	} else if loginCount > 10 {
		result = "watch out"
	} else {
		result = "Excatly 10 login count"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	if num := 3; num < 10 {
		fmt.Println("Number is than 10")
	} else {
		fmt.Println("Number is not less than 10")
	}

}
