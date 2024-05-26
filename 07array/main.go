package main

import (
	"fmt"
)

func main() {
	fmt.Println("Array")
	var fruitList [7]string

	fruitList[0] = "apple"
	fruitList[1] = "kiwi"
	fruitList[2] = "watermelon"
	fruitList[3] = "orange"
	fruitList[4] = "grapes"
	fruitList[5] = "banana"
	fruitList[6] = "apple"

	fmt.Println("Fruitlist is: ", fruitList)
	fmt.Println("Fruitlist is: ", len(fruitList))

	var veglist = [3]string{"potato", "beans", "mushrooms"}
	fmt.Println("Vegy list is ", len(veglist))

}
