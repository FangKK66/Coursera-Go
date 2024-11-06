package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GenDisplaceFn(a float64, v float64, s float64) func(float64) float64 {

	fn := func(t float64) float64 {
		return (0.5*a*t*t + v*t + s)
	}
	return fn
}

func main() {

	var acc float64     // acceleration
	var ini_v float64   // initial velocity
	var ini_dis float64 // initial displacement
	var time float64

	fmt.Printf("Please enter acceleration, initial velocity, and initial displacement: ")
	// Use bufio.NewReader to read user input.
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n') // Read a line

	// Remove whitespace at both ends
	input = strings.TrimSpace(input)

	// Split incoming strings by space
	parts := strings.Split(input, " ")

	acc, _ = strconv.ParseFloat(parts[0], 64)
	ini_v, _ = strconv.ParseFloat(parts[1], 64)
	ini_dis, _ = strconv.ParseFloat(parts[2], 64)

	var inputData string
	fmt.Printf("Please enter time: ")
	fmt.Scan(&inputData)

	time, _ = strconv.ParseFloat(inputData, 64)

	fn := GenDisplaceFn(acc, ini_v, ini_dis)
	fmt.Printf("The final displacement is: %f", fn(time))

}
