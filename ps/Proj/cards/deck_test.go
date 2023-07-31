// Test file study

package main

import (
	"os"
	"testing"
)

// Testing function which are in deck.go and deciding what tests to write

// t is the test handler

// newDeck() - Testing block
func TestNewDeck(t *testing.T) {

	// 1 - Test if the total numbe of cards in the deck are 16 that is the syntax below
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	// 2 - Testing 1st card in dek is the Ace of Spades
	if d[0] != "Ace of ♠️ Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	// 3 - Testing last card is Four ♣️ Clubs
	// Note to access the last element of the slice you need syntax d[len(d)-1] -
	// Special note - These tests are being automaticallyt done by the codeium assandpussy
	if d[len(d)-1] != "Four of ♣️ Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[len(d)-1])
	}
}

// newDeckFromFile() - Testing block
// T is test handler - *testing is ... 
func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {

	// This syntax will remove the testing file which is being created
	os.Remove("_decktesting")

	// THis loads the new into the deck variable and then we are saving it into a file
	deck := newDeck()
	deck.saveToFile("_decktesting")

	// Now loading the deck from the file which has been created
	loadedDeck := newDeckFromFile("_decktesting")

	// As a simple test we are checking the length of the deck loaced from the file
	// Note : We are not checking for the size of the file also when you run the function you done
	// see a file called _decktesting in the directory
	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
