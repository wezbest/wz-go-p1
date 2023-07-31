// 2️⃣Secondary Function

package main

import "fmt"

// New deck type being created for the project
// Create a new type deck
// Which is slice of string
type deck []string

// newDeck returns a new deck of cards.
// This code block is iterating over a list and then  appending each card to the new deck

func newDeck() deck {
	cards := deck{}

	// Defining two lists that has the cards and the value - Note icons were generated from https://www.flaticon.com/
	cardSuits := []string{"♠️ Spades", "♥️ Hearts", "♦️ Diamonds", "♣️ Clubs"}
	cardValue := []string{"Ace", "Two", "Three", "Four"}

	// Iterating over the lists with the _ operator so we dont care about those index values
	for _, suit := range cardSuits {
		for _, value := range cardValue {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards

}

// NOTE : the following function block has
// the received funtion being setup

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
