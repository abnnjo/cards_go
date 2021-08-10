package main

import (
	"testing"
	"os"
)

func TestNewDeck(t *testing.T){
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Spade, but got %v", d[0])
	}

	if d[len(d) - 1] != "Four of Clubs" {
		t.Errorf("Expected last card of Clubs, but got %v", d[0])
	}
}

func TestSaveToDeckNewDeckTestFromFile(t *testing.T){
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadDeck := newDeckFromFile("_decktesting")
	if len(loadDeck) != 16 {
		t.Errorf("expected 16 cards in deck, got %v", len(loadDeck))
	}
	os.Remove("_decktesting")
}