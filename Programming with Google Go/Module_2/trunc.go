package main

import "fmt"

func main() {

	var numFloat float64
	var numInt int64
	fmt.Printf("Please enter a floating point point number: ")
	fmt.Scan(&numFloat)
	numInt = int64(numFloat)
	fmt.Printf("The rounded integer value is: %d\n", numInt)

}
