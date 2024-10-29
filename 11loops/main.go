package main

import "fmt"

func main() {
	fmt.Println("Welcome to Loops")

	//type 1 for "for loop"
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Friday", "Saturday"}
	fmt.Println(days)

	//type 2 for "for loop"
	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	//type 3 for "for loop"
	for i := range days {
		fmt.Println(days[i])
	}

	//type 4 for "for loop"
	for _, day := range days {
		fmt.Println("index is and value is", day)
	}

	//###############################################
	//Example of while loop

	rougueValue := 1

	for rougueValue < 10 {

		if rougueValue == 2 {
			goto soumya
		}

		if rougueValue == 5 {
			rougueValue++
			continue
		}

		fmt.Println("value is: ", rougueValue)
		rougueValue++
	}

soumya:
	fmt.Println("leaning at learncodeonline.in")

	//#########################################################################
	//Example of while loop

	number := 1

	for number <= 10 {
		fmt.Println(number)
		number++
	}

	//#################################################
	//Program to create a multiplication table of 5 using while loop

	multiplier := 1

	//// run while loop for 10 times
	for multiplier <= 10 {

		// find the product
		product := 5 * multiplier

		// print the multiplication table in format 5 * 1 = 5
		fmt.Printf("5 * %d = %d\n", multiplier, product)
		multiplier++
	}
}
