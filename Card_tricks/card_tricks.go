package main

import "fmt"

// As a magician-to-be, Elyse needs to practice some basics. She has a stack of cards that she wants to manipulate.

// To make things a bit easier she only uses the cards 1 to 10.

// Task 1
// Create a slice with certain cards

// When practicing with her cards, Elyse likes to start with her favorite three cards of the deck: 2, 6 and 9. Write a function FavoriteCards that returns a slice with those cards in that order.

// cards := FavoriteCards()
// fmt.Println(cards)
// // Output: [2 6 9]

// Stuck? Reveal Hints
// Opens in a modal
// Task 2
// Retrieve a card from a stack

// Return the card at position index from the given stack.

// card := GetItem([]int{1, 2, 4, 1}, 2) // card == 4
// If the index is out of bounds (ie. if it is negative or after the end of the stack), we want to return -1:

// card := GetItem([]int{1, 2, 4, 1}, 10) // card == -1
// Note
// By convention in Go, an error is returned instead of returning an "out-of-band" value. Here the "out-of-band" value is -1 when a positive integer is expected. When returning an error, it's considered idiomatic to return the zero value with the error. Returning an error with the proper return value will be covered in a future exercise.

// Stuck? Reveal Hints
// Opens in a modal
// Task 3
// Exchange a card in the stack

// Exchange the card at position index with the new card provided and return the adjusted stack. Note that this will modify the input slice which is the expected behavior.

// index := 2
// newCard := 6
// cards := SetItem([]int{1, 2, 4, 1}, index, newCard)
// fmt.Println(cards)
// // Output: [1 2 6 1]
// If the index is out of bounds (ie. if it is negative or after the end of the stack), we want to append the new card to the end of the stack:

// index := -1
// newCard := 6
// cards := SetItem([]int{1, 2, 4, 1}, index, newCard)
// fmt.Println(cards)
// // Output: [1 2 4 1 6]

// Stuck? Reveal Hints
// Opens in a modal
// Task 4
// Add cards to the top of the stack

// Add the card(s) specified in the value parameter at the top of the stack.

// slice := []int{3, 2, 6, 4, 8}
// cards := PrependItems(slice, 5, 1)
// fmt.Println(cards)
// // Output: [5 1 3 2 6 4 8]
// If no argument is given for the value parameter, then the result equals the original slice.

// slice := []int{3, 2, 6, 4, 8}
// cards := PrependItems(slice)
// fmt.Println(cards)
// // Output: [3 2 6 4 8]

// Stuck? Reveal Hints
// Opens in a modal
// Task 5
// Remove a card from the stack

// Remove the card at position index from the stack and return the stack. Note that this may modify the input slice which is ok.

// cards := RemoveItem([]int{3, 2, 6, 4, 8}, 2)
// fmt.Println(cards)
// // Output: [3 2 4 8]
// If the index is out of bounds (ie. if it is negative or after the end of the stack), we want to leave the stack unchanged:

// cards := RemoveItem([]int{3, 2, 6, 4, 8}, 11)
// fmt.Println(cards)
// // Output: [3 2 6 4 8]

func FavoriteCards() []int {
	cards := []int{2,6,9}
	return cards
}

func GetItem(slice []int, index int) int{
	if index <0 || index >= len(slice) {
		return -1
	} else {
		return slice[index]
	}

}

func SetItem(slice []int, index int, newCard int) []int{
	if index < 0 || index >= len(slice){
		return append(slice, newCard)
	} else {
		slice[index] = newCard
		return slice
	}

}


func PrependItems(slice []int, value ...int)[]int{
	return append(value, slice...)
}


func RemoveItem(slice []int, index int) []int {
	if index < 0 || index >= len(slice){
		return slice
	} else {
		return append(slice[:index], slice[index+1:]...)
	}
}

func main(){
	fmt.Println(FavoriteCards())
	fmt.Println(GetItem([]int{1,2,4,1},2))
	fmt.Println(SetItem([]int{1,2,4,1},10,6))
	fmt.Println(PrependItems([]int{3,2,6,4,8},5,1))
	fmt.Println(RemoveItem([]int{3,2,6,4,8},2))
}