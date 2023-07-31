// 🕐 Primary Function

package main

func main() {
	// cards := newDeck()
	// cards.saveToFile("mistress")

	// cards := newDeckFromFile("mistress")
	// cards.print()

	// Using newdeck() to store slice into variable cards
	cards := newDeck()

	// This function uses the randomizer
	cards.shuffle()

	// This is the print function defined in deck.go
	cards.print()

	//This function written by you , to save the date which is being used as the seed in the
	// randomizer function
	cards.saveToFilewDate("panty.txt")
}
