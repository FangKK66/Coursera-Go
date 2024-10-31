package main

import (
	"fmt"
	"sort"
	"strconv"
	"unicode"
)

func isNumeric(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

func main() {

	var inputData string
	var numInt int
	var exitNow int = 0
	sli := make([]int, 0, 100)

	for exitNow != 1 {
		fmt.Printf("Please enter a integer number: ")
		fmt.Scan(&inputData)
		if inputData == "X" || inputData == "x" {
			exitNow = 1
			break
		}
		if !isNumeric(inputData) {
			fmt.Printf("Wrong input type, integer only!!!\n")
			continue
		}
		numInt, _ = strconv.Atoi(inputData)
		sli = append(sli, numInt)
		sort.Ints(sli)
		index := 1
		for index < len(sli)+1 {
			fmt.Printf("%d. %d\n", index, sli[index-1])
			index = index + 1
		}
	}

}
