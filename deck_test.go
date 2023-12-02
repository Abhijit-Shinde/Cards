package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 165, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected First card to be Ace of Spades but got %v", d[2])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card to be Four of Clubs but got %v", d[15])
	}

}

func TestSaveToFile_DeckFromFile(t *testing.T) {
	os.Remove("deck_testing.txt")

	deck := newDeck()
	deck.saveToFile("deck_testing.txt")

	readFile := deckFromFile("deck_testing.txt")
	if len(readFile) != 16 {
		t.Errorf("Expected 18 cards but got %v", len(readFile))
	}

	os.Remove("deck_testing.txt")
}
