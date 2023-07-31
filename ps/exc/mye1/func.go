package main

import "fmt"

func MyFunc() {

	// create integer slice

	// num := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	numo := MakeSlice()
	// This is the for function
	// This solution which was given is not completely accurate
	// So don rely on it too muhc , also note that in the tut
	// Solution , they used the same function which was used to make the original deck
	// That should hav been used instead
	for i := 0; i < len(numo); i++ {
		if numo[i]%2 == 0 {
			fmt.Println(numo[i], "is even")
		} else {
			fmt.Println(numo[i], "is odd")
		}
	}

}

// Making a seperate function to make the slice for testing
func MakeSlice() []int {
	num := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	return num
}

// This solution is automated
func AuGe() {
	var numbers []int = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < len(numbers); i++ {
		if numbers[i]%2 == 0 {
			fmt.Println(numbers[i], "is even")
		}
	}
}
