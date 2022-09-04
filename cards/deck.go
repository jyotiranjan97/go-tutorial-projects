package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

// d <- The actual copy of the deck we're working with
//		is available in the function as a variable 'd'.(Receiver)
//		Should be a single letter abbreviation
// deck <- Every variable of type 'deck'
//			can call this function on itself
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
