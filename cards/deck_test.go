package main

import "testing"

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
