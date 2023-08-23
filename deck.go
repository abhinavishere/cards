package main

import "fmt"

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := deck{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := deck{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	// Iterating over array of cards
	for _, card := range d {
		fmt.Println(card)
	}
}

func (d deck) printLength() {
	fmt.Println(len(d))
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
