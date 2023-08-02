package main

import "fmt"

func Rang() {
	fruits := []string{"apple", "banana", "cherry"}
	cars := []string{"Ford", "Volvo", "BMW"}

	for i, fruit := range fruits {
		fmt.Println(i, fruit)
	}

	for i, car := range cars {
		fmt.Println(i, car)
	}
}
