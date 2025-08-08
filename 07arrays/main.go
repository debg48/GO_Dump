package main

import "fmt"

func main() {

	fmt.Println("Welcome to Arrays Playground")

	var arr [4]string
	arr[0] = "apple"
	arr[1] = "banana"
	// arr[2] = "orange"
	arr[3] = "grapes"

	fmt.Println("The fruit list is : ", arr)
	fmt.Println("The length of array is : ", len(arr))

	var vegList = [3]string{"potato", "beans", "mushroom"}
	fmt.Println("The veg list is : ", vegList)
}
