package main

import "fmt"

var deckSize int

func main() {
	// var card = "Ace of spades"
	// Can only be used to assign a new variable not to reassign
	card := "Ace of spades"
	fmt.Println(card)
	card = "Five of diamonds"
	fmt.Println(card)

	deckSize = 52
	fmt.Println(deckSize)
}
