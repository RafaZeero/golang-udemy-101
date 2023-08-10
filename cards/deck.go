package main

import "fmt"

// Create a new type of "decks"
// which is a slice of strings
type deck []string

func newCard(rank string, suit string) string {
	newCard := fmt.Sprintf("%s of %s", rank, suit)
	return newCard
}

func newDeck() deck {
	cards := deck{}
	ranks := []string{"Ace", "Two", "Three", "Four"}
	suits := []string{"Heart", "Diamond", "Club", "Spade"}

	for _, suit := range suits {
		for _, rank := range ranks {
			cards = append(cards, newCard(rank, suit))
		}
	}

	return cards

}

func (d deck) print() {
	for idx, card := range d {
		fmt.Println(idx, card)
	}
}