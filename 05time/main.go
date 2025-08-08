package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of GO Lang")
	// currentTime := time.Now()
	// fmt.Println(currentTime)
	// fmt.Println(currentTime.Format("01-02-2006 Monday 15:04:05"))

	createDate := time.Date(2024,time.September,30,23,23,0,0,time.UTC)
	fmt.Println(createDate)
	fmt.Println(createDate.Format("01-02-2006 Monday"))
}
