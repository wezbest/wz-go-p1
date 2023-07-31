package main

import (
	f "fmt"
)

// here defining the lipgloss colors

func gotch() {
	mySlice := []string{"hi", "how", "are", "you", "doing"}

	upGoc(mySlice)
	f.Println("\n\n\n")
	colo("Printing of gotchas")
	f.Println(mySlice)

}

// Slice works differetnt thatn sructs , you dont need

func upGoc(s []string) {
	s[0] = "ChangeHi"
 
}
