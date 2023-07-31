/*
Actual solution of the problem
*/

package pkg

import (
	co "ex3/pkg/config" // This adds coloring
	"fmt"
	"io"
	"log"
	"os"
)

// Defining interface to prevent circular dependencies issues
type Trie interface {
	Tree()
}

func Tree() {
	co.Texc(`
Actual Folder structure
-----------------------
├── cmd
│  └── main.go
├── go.mod
├── go.sum
└── pkg
   ├── config
   │  └── lip.go
   └── work
	`, "")
}

// This is for openign the file
func OpenFile() {

	co.Headr(`
My Solution  - My Solution - My Solution - My Solution 	`)

	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <input>")
		os.Exit(1)
	}

	// Get input argument
	input := os.Args[1]

	f, err := os.OpenFile(input, os.O_RDWR, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err = io.Copy(os.Stdout, f)
	if err != nil {
		log.Fatal(err)
	}

}
