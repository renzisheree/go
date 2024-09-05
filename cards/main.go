package main

func main() {

	// hand , remainingCard :=deal(cards,5)
	// hand.print()

	// remainingCard.print()
	// cards := newDeck();
	// cards.saveToFile("myCard")
	// cards := readFile("myCard1")
	cards:= newDeck()
	cards.print()
	cards.shuffle()
	cards.print()
}
