package main

import "fmt"

func main() {
	//var card string = "Ace os Spaded"
	card := "Ace of Spades" // only for defining
	// card := newCard()
	fmt.Println(card)
	//fmt.Println("hey")

	cards := deck{"xyz", newCard(), "zzz"}

	cards = append(cards, "DU")
	//cards.print()

	//hand, remaining := deal(cards, 2)
	// hand.print()
	//fmt.Println("fsdbf")
	//remaining.print()
	//cards.saveToFile("myfile")
	newcards := newDeck("myfile")
	//newcards.print()
	newcards.shuffle()
	newcards.print()
	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }
}

func newCard() string {
	return "5 of Diamonds"
}
