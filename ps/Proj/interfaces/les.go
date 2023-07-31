/*
Actual lessons code all go in here
- Writing the bots progrm hers as two seperae functions here where we will call it in main.gol
*/

package main

import "fmt"

// ++++++++++++++++++++++++++++ Interface definitions

type bot interface {
	getGreeting() string
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++

// Struc interface definitions
type englishBot struct{}
type bootishBot struct{}

/*
English Both Code here
*/

func (englishBot) getGreeting() string {

	// VERY Custom Logic for generating Enligh Greeting

	return "Lick Her Now"
}

// English Greeting

// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }

/*
Bootish Both Code here
*/

func (bootishBot) getGreeting() string {

	// VERY Custom Logic for generating Bootish Greeting

	return "Sniffff"
}

// func printGreeting(bb bootishBot) {
// 	fmt.Println(bb.getGreeting())
// }
