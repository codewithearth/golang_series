package main

import "fmt"

const LoginToken string = "dgshrgshvas" //Public

func main() {
	fmt.Println("Variables")

	//string
	var username string = "Prince"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	//boolien
	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	//int
	var smallval uint8 = 255
	fmt.Println(smallval)
	fmt.Printf("Variable is of type: %T \n", smallval)

	//float
	var smallFloat float32 = 255.234235232
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	//default values and some alises
	//int
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	//string
	var anotherStringVaribale string
	fmt.Println(anotherStringVaribale)
	fmt.Printf("Variable is of type: %T \n", anotherStringVaribale)

	//implicit type
	var website = "codewithearth.in"
	fmt.Println(website)

	//no var style
	numberOfUser := 3000000
	fmt.Println(numberOfUser)

	//const
	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)
}
