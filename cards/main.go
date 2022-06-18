package main

import (
	"fmt"
	"strings"
)

// var card string

func main() {
	// Can initilize the value out of function and set value
	// card = "Ace of Spades"

	// var card string
	// card = "Ace of Spades"

	// var card string = "Ace of Spades"

	// Does not add but creates new slice
	// append()

	cards := newDeck()
	hand, remainingDeck := deal(cards, 5)
	hand.print()
	fmt.Println("--")
	remainingDeck.print()
	fmt.Println("--")
	fmt.Println(cards)
	fmt.Println([]string(cards))
	fmt.Println("--")
	fmt.Println(strings.Join(cards, ","))
	fmt.Println(cards.toString())
	fmt.Println("--")
	cards.saveToFile("my_cards")
	newCards := newDeckFromFile("my_cards")
	newCards.print()
	fmt.Println("--")
	newCards.shuffle()
	newCards.print()

}
