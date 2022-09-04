package main

//* var card string = "Ace of Spades" // can execute
//* var card string // can execute
//! card := "Ace of Spades" // run time error, := is not allowed outside func

func main() {

	// var <- We're about to create a  new variable
	// card <- The name of variable will be 'card'
	// string <- The type of variable is 'string'
	// = <- Assignment operator
	// 'Ace of Spaded' <- This value is assigned to the variable
	//* var card string = "Ace of Spades"

	// := <- This is to be used only during initialization of variable
	//* card := "Ace of Spades"
	//* card = "Five of Diamonds"

	//* card := newCard()

	//* fmt.Println(card)

	// Slice
	// - Like a dynamic array
	// - Must contain same data types
	// - syntax, []#dataType{#data} // # is a placeholder
	//* cards := []string{"Ace of Diamonds", newCard()}

	//* cards := deck{"Ace of Diamonds", newCard()}
	//* cards = append(cards, "Six of Spaded")

	cards := newDeck()

	// range cards is returning two things, 1st is index & 2nd is the value
	// used as an iterator for slice
	//* for i, card := range cards {
	//* 	fmt.Println(i, card)
	//* }

	cards.print()
}

// newCard <- Define a function called 'newCard'
// string <- When executed, this function will return a value of type 'string'
//! Without specifying return type(string), compiler will through an error,
//! too many return statements, as it thinks nothing will return
//* func newCard() string {
//* 	return "Five of Diamonds"
//* }
