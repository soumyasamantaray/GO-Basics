package main

import "fmt"

func main() {
	fmt.Println("Welcome to the Arraylist")

	//Defining the Arraylist

	var fruitList [5]string

	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[2] = "Grape"
	fruitList[3] = "Peach"
	fruitList[4] = "Cocuanut"

	fmt.Println("Fruit lists are : ", fruitList)
	fmt.Println("Fruit leanth is : ", len(fruitList))

	//ArrayList in single line

	var vegList = [4]string{"Patato", "tamaato", "Beans", "Gralic"}
	fmt.Println("Veg lists are : ", vegList, "& Number : ", len(vegList))

}
