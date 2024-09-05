package main

import (
	"os"
	"testing"
)


func TestNewDeck(t *testing.T) {
	d := newDeck()
	if(len(d) != 16) {
		t.Errorf("Expected deck length of 16 but got %d" , len(d))
	}
}

func TestDeckValue(t *testing.T) {
	d := newDeck()
	if(d[0] != "Ace of Diamonds") {
		t.Errorf("Expected card to be Ace of Spades but got %v" , d[0])
	}
	if d[len(d) -1] != "Four of Clubs" {
		t.Errorf("Expected card to be Four of Clubs but got %v" , d[len(d) -1])
	}
}

func  TestIO(t *testing.T) {
	os.Remove("_decktesting")
	d := newDeck()
	d.saveToFile("_decktesting")
	

	loadedDeck := readFile("_decktesting")

	if(len(d) != 16) {
		t.Errorf("Expecting 16 cards in deck but got %v" , len(loadedDeck))
	}	
	os.Remove("_decktesting")

}