package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name string `json:"name"`
	Addr string `json:"addr"`
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var inputName string
	var inputAddr string

	fmt.Printf("Please enter your name: ")
	inputName, _ = reader.ReadString('\n')
	inputName = inputName[:len(inputName)-1]
	// fmt.Scanln(&inputName) // “Scanln” will read the entire line of input including spaces
	// fmt.Scan(&inputName)
	fmt.Printf("Please enter your address: ")
	inputAddr, _ = reader.ReadString('\n')
	inputAddr = inputAddr[:len(inputAddr)-1]
	// fmt.Scanln(&inputAddr)
	// fmt.Scan(&inputAddr)

	p1 := Person{Name: inputName, Addr: inputAddr}

	barr, err := json.Marshal(p1)
	if err != nil {
		fmt.Println("Error serializing to JSON:", err)
		return
	}

	fmt.Println(string(barr))

}
