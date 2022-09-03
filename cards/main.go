package main

import "fmt"

//* var card string = "Ace of Spades" // can execute
//* var card string // can execute
//! card := "Ace of Spades" // run time error, := is not allowed outside func

func main() {

	// var <- We're about to create anew variable
	// card <- The name of variable will be 'card'
	// string <- The type of variable is 'string'
	// = <- Assignment operator
	// 'Ace of Spaded' <- This value is assigned to the variable
	//* var card string = "Ace of Spades"

	// := <- This is to be used only during initialization of variable
	card := "Ace of Spades"
	card = "Five of Diamonds"

	fmt.Println(card)
}
