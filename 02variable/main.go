package main

import "fmt"

const LoginToken string = "testingthetoken" //In go if a variable is starting with capital letter then it became capital

func main() {
	var test string = "hemant"
	var isLoggedIn bool = true
	var testinteger int = 256
	var testunsignedinteger uint8 = 255
	var testingfloat float64 = 255.6543224
	fmt.Println("username", test)
	fmt.Printf("variable is of type : %T \n", test)
	fmt.Println("boolen test", isLoggedIn)
	fmt.Printf("checking the type of boolen data type : %T \n", isLoggedIn)
	fmt.Println("integer", testinteger)
	fmt.Printf("test the type of testinteger : %T \n", testinteger)
	fmt.Println("the value for unsigned integer", testunsignedinteger)
	fmt.Printf("the type of testunisgnedinteger is %T \n", testunsignedinteger)
	fmt.Println("testing the float", testingfloat)
	fmt.Printf("testing the type of float type %T \n", testingfloat)

	// default values and some aliases
	var anothervariable int
	fmt.Println(anothervariable)
	fmt.Printf("variable is of type : %T \n", anothervariable)
	//implict type or way to declare the variable in go
	var website = "hemantkumar"
	fmt.Println(website)
	// no var style
	numberofuser := 23 //remember the colon before equal is more important to declare the variable without that it will show an error
	fmt.Println(numberofuser)
	fmt.Println(LoginToken)
	fmt.Printf("the type of public variable %T \n", LoginToken)
}
