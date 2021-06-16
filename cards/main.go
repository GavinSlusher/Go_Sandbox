package main

func main() {

	cards := newDeck()
	cards.shuffle()

	// cards.saveToFile("my_cards")

	cards.print()
}

// Cards OO Approach
// 	- Create a Deck Class
// 		- Creates a Deck Instance
// 		- Print function
// 		- Shuffle function
// 		- SaveToFile function

// Go Approach
// 	- Create a type
// 		- type deck []string
// 			- telling go we want to create an array of strings
// 			- make functions specifically to work with this array
// 	- Create functions with 'deck' as a 'receiver'
// 		- like a "method" - a function that belongs to an instance

// Project Structure
// 	- main.go
// 		- code to create and manip the deck
// 	- deck.go
// 		- code that describes what a deck is and how it works
// 	-deck_test.go
// 		- code to automatically test the deck
