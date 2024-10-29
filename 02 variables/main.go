package main

import "fmt"

func main() {

	var print string = "Hi"
	fmt.Println(print)
	fmt.Printf("this is the variable file : %T \n", print)

	var Isbool bool = true
	fmt.Println(Isbool)
	fmt.Printf("this is the variable file : %T \n", Isbool)

	var Isunit uint8 = 255
	fmt.Println(Isunit)
	fmt.Printf("this is the variable file : %T \n", Isunit)

}
