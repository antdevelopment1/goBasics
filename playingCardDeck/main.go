package main

import "fmt"

func main() {
	// var card = "Ace of spades"
	// Can only be used to assign a new variable not to reassign
	cards := []string{"5 of spades", newCard()}
	cards = append(cards, "Jack of diamonds")
	fmt.Println(cards)

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

// We have to tell this function what data type it should be expecting to return any time this function is called
func newCard() string {
	return "Five of diamonds"
}
