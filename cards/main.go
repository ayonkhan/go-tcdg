package main

import "fmt"

func main() {
	//cards := newDeck()
	//
	//hand, remainingCards := deal(cards, 5)
	//
	//hand.print()
	//remainingCards.print()

	//cards := newDeck()
	//cards.saveToFile("my_cards.txt")

	cards := newDeck()
	cards.shuffle()
	fmt.Println(cards)
}
