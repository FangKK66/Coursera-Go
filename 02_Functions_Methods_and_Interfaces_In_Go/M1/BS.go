package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Swap(sli *[]int, idx int) {
	temp := (*sli)[idx]
	(*sli)[idx] = (*sli)[idx+1]
	(*sli)[idx+1] = temp
}

func BubbleSort(sli *[]int) {
	n := len(*sli)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			// If the former element is larger than the latter, exchange them
			if (*sli)[j] > (*sli)[j+1] {
				Swap(sli, j)
			}
		}
	}
}

func main() {
	fmt.Printf("Please enter 10 numbers(using space between numbers): ")

	// Use bufio.NewReader to read user input.
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n') // Read a line

	// Remove whitespace at both ends
	input = strings.TrimSpace(input)

	// Split incoming strings by space
	parts := strings.Split(input, " ")

	// Check if ten integers were entered
	if len(parts) != 10 {
		fmt.Println("Please enter ten integer numbers.")
		return
	}

	// Converting strings to integers
	numberList := make([]int, 10)
	for i, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			fmt.Println("Invalid entry, please enter a valid integer.")
			return
		}
		numberList[i] = num
	}

	fmt.Printf("\nGiven input: %v", numberList)

	BubbleSort(&numberList)

	fmt.Printf("\nAfter BSort: %v\n", numberList)

	// index := 0
	// for index < len(numberList) {
	// 	fmt.Printf("%d\n", numberList[index])
	// 	index = index + 1
	// }
}
