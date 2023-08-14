package main

func main() {
	// deck := newDeck()
	// deck.saveToFile("cartinhas")

	deck := newDeckFromFile("cartinhas")

	deck.print()

	return
}
