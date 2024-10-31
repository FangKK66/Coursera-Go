package main

import (
	"fmt"
	"os"
	"strings"
)

type Name struct {
	Fname string `json:"first name"`
	Lname string `json:"last name"`
}

func main() {

	var fileName string
	var firstName string
	var lastName string
	var nameBox Name

	fmt.Printf("Please enter file name: ")
	fmt.Scan(&fileName)
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error file name:", err)
		return
	}
	barr := make([]byte, 10000)
	nb, err := f.Read(barr)
	f.Close()

	content := string(barr[:nb])
	words := strings.Fields(content)

	sli := make([]Name, 0, 100)

	for idx, word := range words {
		if idx%2 == 0 {
			firstName = word
		} else {
			lastName = word
			nameBox.Fname = firstName
			nameBox.Lname = lastName
			sli = append(sli, nameBox)
		}
	}

	//

	index := 0
	for index < len(sli) {
		fmt.Printf("%s\n", sli[index])
		index = index + 1
	}

}
