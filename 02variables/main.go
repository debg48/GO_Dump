package main

import "fmt"

const loginToken string = "ABCD"

func main() {
	var username string = "Debgandhar"
	fmt.Println(username)
	fmt.Printf("Variable is of type : %T\n", username)

	var IsLoggedIn bool = true
	fmt.Println(IsLoggedIn)
	fmt.Printf("Variable is of type : %T\n", IsLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type : %T\n",smallVal)

	var smallFloat float32 = 255.456789456789
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type : %T\n",smallFloat)

	//Aliases and Defaults
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type : %T\n",anotherVariable)

	var website = "www.google.com"
	fmt.Println(website)

	noOfUser := 3000
	fmt.Println(noOfUser)

	fmt.Println(loginToken)

}
