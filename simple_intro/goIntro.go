package main

import "fmt"

func main() {
	// create a new var
	// name of the var will be 'card'
	// only a 'string' will ever be assigned to this var
	// var card string = "Ace of Spades"
	// might be better to:
	// card := "Ace of Spades"
	// card = "Five of Diamonds"
	// Go compiler will figure our that card is a string
	// Note: only use `:=` when first declaring a var. If reassigning the var,
	// `:=` is not needed

	cards := []string{"Ace of Diamonds", newCard()}

	// note, this returns a NEW slice w/the new element
	cards = append(cards, "Six of Spades")

	// interate over the index element, `i`, using the current element `card`
	// by taking the slice of `cards` and looping over it
	// using the := because we're technically 'redeclaring' `i` and `card`
	// each time we iterate over the loop
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}

// // Two different list-like structures for Go:

// 1) Array
// - A basic list of records
// - fixed length

// 2) Slice
// - more feature-rich list
// - can grow or shrink

// 3) Both
// - must be declared to be of the same type. (all elements must be the same)
