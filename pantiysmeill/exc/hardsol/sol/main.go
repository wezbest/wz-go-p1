/*
Main solution file
*/

package main

import (
	"har/files"
	"har/src"
	// "har/files"
)

func main() {
	src.ColoBanner("Main Solution")

	src.Headr("Solution")
	files.Solu()

	// src.Headr("Printing curl command here")
	// files.Solu2()

	src.Headr("This is piping commands")
	files.Solu3()
}
