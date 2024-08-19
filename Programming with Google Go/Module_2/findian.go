package main

import (
	"fmt"
	"strings"
)

func main() {

	var enterR string
	var lowerR string
	fmt.Printf("Please enter a string: ")
	fmt.Scan(&enterR)
	lowerR = strings.ToLower(enterR)
	var hasA bool = strings.Contains(lowerR, "a")
	var preI bool = strings.HasPrefix(lowerR, "i")
	var sufN bool = strings.HasSuffix(lowerR, "n")
	if hasA && preI && sufN {
		fmt.Printf("Found!\n")
	} else {
		fmt.Printf("Not Found!\n")
	}

}
