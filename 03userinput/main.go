package main

import (
	"bufio"
	"fmt"
	"os"
)

//Basic codes for user inputs

func main() {

	var welcome string = "basic codes"
	fmt.Println("Welcome to user input", welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Pizza ")

	//
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for ratings", input)
	fmt.Printf("Type of this rating is %T", input)

}
