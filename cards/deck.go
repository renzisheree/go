package main

import "fmt"

//create new type of deck
// which is a slice of string

type deck []string


func newDeck() deck {
 cards:= deck{} 
 cardSuits := []string{"Diamonds" , "Hearts" , "Spades" , "Clubs"}
 cardValues := []string{"Ace" , "Two" , "Three" , "Four"}
 for _, suit := range cardSuits {
	for _ , value := range cardValues {
		cards = append(cards, value + " of " + suit)
	}
 }
 return cards
}

func (d deck) print() {
	for i , card := range d {
		fmt.Println(i,card)
	}}

func deal(d deck , handSize int) ( deck, deck) {
	return d[:handSize] , d[handSize:]
} 