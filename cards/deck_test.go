package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	/* Test deck size */
	length := 16
	if len(d) != length {
		t.Errorf("Exoected deck length to be %v, but instead got %v", length, len(d))
	}

	/* Test deck first card */
	firstCard := "Ace of Hearts"
	if d[0] != firstCard {
		t.Errorf("Expected deck first card to be %s, but instead got %s", firstCard, d[0])
	}

	/* Test deck last card */
	lastCard := "Four of Spades"
	if d[len(d)-1] != lastCard {
		t.Errorf("Expected deck last card to be %s, but instead got %s", lastCard, d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	fileName := "_decktesting"
	os.Remove(fileName)

	deck := newDeck()
	deck.saveToFile(fileName)

	loadedDeck := newDeckFromFile(fileName)

	/* Test deck and loaded deck sizes match */
	if len(loadedDeck) != len(deck) {
		t.Errorf("Expected loadedDeck length to be %v, but instead got %v", len(deck), len(loadedDeck))
	}

	/* Test loaded deck sizes equals total deck size */
	totalSize := 16
	if len(loadedDeck) != totalSize {
		t.Errorf("Expected loadedDeck length to be %v, but instead got %v", totalSize, len(loadedDeck))
	}

	/* Cleanup files */
	defer os.Remove(fileName)

}
