package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

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

func (d deck) toString() string {
	return strings.Join([]string(d) , ",")
}

func (d deck) saveToFile(fileName string) error {
	return os.WriteFile(fileName,[]byte(d.toString()), 0666)
}
func readFile(fileName string) (deck) {
	bs,err := os.ReadFile(fileName)
	if err != nil  {
		fmt.Println("Error",err)
		os.Exit(1)
	}
	s :=strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	for index  := range d {
		newPosition := rand.Intn(len(d) -1)
		d[index], d[newPosition] = d[newPosition] , d[index]
	}
}