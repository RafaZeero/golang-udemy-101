package main

import (
	"fmt"
	"strings"
)

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

func deal(d deck, handSize int) (deck, deck) {
	/** Returns TOTAL cards for handsize, REMAINING cards */
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	res := strings.Join([]string(d), ", ")
	fmt.Println(res)
	return res
}
