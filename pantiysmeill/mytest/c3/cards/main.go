// 🕐 Primary Function

package main

func main() {
	// cards := newDeck()
	// cards.saveToFile("mistress")

	cards := newDeckFromFile("mistress")
	cards.print()
}
