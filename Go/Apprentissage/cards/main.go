package main

func main() {
	cards := newDeck()

	//hand, remainingCards := deal(cards, 5)
	//fmt.Println("Hand")
	//hand.print()
	//fmt.Println("Remaining Cards")
	//remainingCards.print()

	cards.saveToFile("myCards")
	cards2 := newDeckFromFile("myCards")
	cards2.shuffle()
	cards2.print()
}
