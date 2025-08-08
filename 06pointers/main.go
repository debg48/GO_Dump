package main

import "fmt"

func main() {
	fmt.Println("Welcome to Pointer Playground")
	// var ptr *int
	// fmt.Println("Value of Pointer is ",ptr)
	// fmt.Printf("Type of Pointer is : %T\n",ptr)
	myNumber := 23
	var ptr  = &myNumber
	fmt.Println("Value of Pointer is ",ptr)
	fmt.Println("Value of Pointer is ",*ptr)
	fmt.Printf("Type of Pointer is : %T\n",ptr)

	*ptr = *ptr + 2
	fmt.Println("The new value of variable is : ",myNumber)

}
