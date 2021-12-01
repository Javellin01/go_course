package main

import "fmt"

type deck []string

func newDeck() deck {
	deck := deck{}

	suites := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	values := []string{"Two", "Three", "Four", "Ace"}

	for _, suite := range suites {
		for _, value := range values {
			deck = append(deck, value+" of "+suite)
		}
	}

	return deck
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}
