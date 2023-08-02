// 2️⃣Secondary Function - All main functions for this project are in here

package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// New deck type being created for the project
// Create a new type deck
// Which is slice of string
// Custom type
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
// this enables d.print() adding the .print() function , you need to have a receiver to do this

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Deal function where with two paratmeters and 2 outputs
// This type is called a slice which has special functions
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// This is the function required for the following
// slice -> string -> byteSlice - Not doing directly  since we might beed this string interim ste
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// Actual writing to the file
func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)

}

// Function to write to a file with current time appended to it
// NOTE : This function has been written by your and not mentioned in the main coure
func (d deck) saveToFilewDate(filename string) error {
	ct := fmt.Sprint(ntime())
	str := d.toString() + "\n SmellherAss :   " + ct
	return os.WriteFile(filename, []byte(str), 0666)

}

// FIle reader - Now the opposite
// []Byte -> string

func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)

	// Error handling block
	if err != nil {
		fmt.Println("😡Error: ", err)
		os.Exit(1) // Bunch of exit codes
	}

	s := strings.Split(string(bs), ",")
	return deck(s)

}

// Shuffler

func (d deck) shuffle() {

	// Setting the seed as the current time in nanoseconds from unix time , this is
	// Part of the standard library
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

// This is your own own function to see the nano time

func ntime() int64 {
	t := time.Now().UnixNano()
	fmt.Println("\n\nCurrent Unix time: ", t)
	return t
}
