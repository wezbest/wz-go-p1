/*
Excercise
1. Read a file and dump it to the terminal
2. Accept file as an inpur parameter
3. Actual work in wo.go
*/

package wo

import (
	"fmt"
	"os"
	"rfd/lip"
)

//  ++ MY WORKS SECTION ROM READING

// +++++++++++++++++++++++++++++++++++++++

// ++++ THIS SECTION WAS A SUMP FROM CALUDE
func ClaudeprintFile(filename string) {

	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	fmt.Print(string(data))
}

func ClaudeMainReadFunc() {

	lip.Hea("This function will print contents of file to screen")

	if len(os.Args) != 2 {
		fmt.Println("😡  😡 😡  😡 😡  😡")
		fmt.Println("Usage: printfile <filename>")
		fmt.Println("😡  😡 😡  😡 😡  😡")
		os.Exit(1)
	}

	filename := os.Args[1]

	filo := string(filename)

	lip.Texc("Filename chosen: ", filo)
	ClaudeprintFile(filename)
}

// +++++++++++++++++++++++++++++++++++++++++++++++++++++
