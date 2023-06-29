package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("La taille doit être de 52 cartes mais on a %v", len(d))
	}
	if d[0] != "Ace of Spades" {
		t.Errorf("The first card should be a Ace of Spades but we got a %v", d[0])
	}
	if d[len(d)-1] != "Valet of Clubs" {
		t.Errorf("The first card should be a Ace of Spades but we got a %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Pas copié proprement car la taille est de %v", len(deck))
	}
	os.Remove("_decktesting")

}
