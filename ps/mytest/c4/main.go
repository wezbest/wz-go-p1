// 🕐 Primary Function

package main

func main() {
	// cards := newDeck()
	// cards.saveToFile("mistress")

	// cards := newDeckFromFile("mistress")
	// cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()
	cards.saveToFilewDate("cjito")
	ntime()
}
