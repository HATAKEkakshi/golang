package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var welcome string = "Welcome to user input"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the password for attack:")

	//comma ok syntax || err ok

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for the password", input)
	fmt.Printf("Type of the password : %T \n", input)
	username()
	inputtesting()
	marks()
}
func inputtesting() {
	password := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the new password")
	newinput, _ := password.ReadString('\n')
	fmt.Println("Here is new password", newinput)
}
func username() {
	user := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the username :")
	userinput, _ := user.ReadString('\n')
	fmt.Println("The Username is :", userinput)
}
func marks() {
	marksofstudent := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the marks of the student")
	marksinput, _ := marksofstudent.ReadString('\n')
	fmt.Println("the marks of student is :", marksinput)
}
