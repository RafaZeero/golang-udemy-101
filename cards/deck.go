package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
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
	suits := []string{"Hearts", "Diamonds", "Clubs", "Spades"}

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
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	return deck(strings.Split(string(data), ","))
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())

	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d))

		d[i], d[newPosition] = d[newPosition], d[i]

	}
}
