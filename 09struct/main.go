package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to struct")

	//using struct

	Ram := User{"Ram", "ram@yahoo.com", true, 16}
	fmt.Println(Ram)
	fmt.Printf("Name is %v and Email is %v.", Ram.Name, Ram.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
