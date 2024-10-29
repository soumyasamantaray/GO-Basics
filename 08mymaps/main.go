package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to maps in golang")

	//adding maps

	languages := make(map[string]string)

	languages["js"] = ("javascript")
	languages["py"] = ("python")
	languages["rb"] = ("ruby")

	fmt.Println("these are our languages : ", languages)
	fmt.Println("js value for  : ", languages["js"])

	delete(languages, "rb")
	fmt.Println("list of all languages:", languages)

	//using for loop

	for key, value := range languages {
		fmt.Printf("for key %v, valueis %v\n", key, value)
	}

	//replacing "key" with "_"

	for _, value := range languages {
		fmt.Printf("for key v, value is %v\n", value)
	}

}
