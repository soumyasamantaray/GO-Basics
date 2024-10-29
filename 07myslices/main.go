package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slices") //Initial print

	//List with single line variable

	var vegList = []string{"Banana", "mango", "grapes"}
	fmt.Println("Thesse are the fruits list : ", vegList)

	//Veg list with Slaices

	vegList = append(vegList, "patato", "tamato", "beans")
	fmt.Println(vegList)

	vegList = append(vegList[1:3])
	fmt.Println(vegList)

	// append numeric value by make ()"append()" and sort.int()

	highScore := make([]int, 4) //appying slice

	highScore[0] = 145
	highScore[1] = 245
	highScore[2] = 345
	highScore[3] = 445

	highScore = append(highScore, 545, 645, 745) //append more value

	sort.Ints(highScore) //other way with append
	fmt.Println(highScore)
	fmt.Println(sort.IntsAreSorted(highScore)) //output boolean

	//substract the value by using "append" and "index"

	var courses = []string{"python,", "go,", "java,", "ruby,"}
	fmt.Println(courses)
	var index int = 2 //"index" substring in the original string
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
