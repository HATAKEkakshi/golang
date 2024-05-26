package main

import "fmt"

func main() {
	fmt.Println("Pointers")
	//var ptr *int while simple declaring a pointer * can be used
	// fmt.Println("value of pointer is ", ptr)
	mynumber := 23

	var ptr = &mynumber // for referceing & is used
	*ptr = *ptr * 2
	fmt.Println("value of ptr", ptr)
	fmt.Println("value of mynumber", mynumber)
	fmt.Println("value of actual pointer", *ptr)
	fmt.Printf("the type of this pointer %T \n", ptr)
}
