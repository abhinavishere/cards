package main

import "fmt"

func main() {
	cards := newDeck()

	// cards.print()
	cards.printLength()

	hand, _ := deal(cards, 5)
	fmt.Println(hand)
}
