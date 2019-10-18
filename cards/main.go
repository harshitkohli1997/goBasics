package main

func main() {
	//var card string = "Ace os Spaded"
	//card := "Ace of Spades" // only for defining
	// card := newCard()
	// fmt.Println(card)
	//fmt.Println("hey")

	cards := deck{"HEY", newCard(), "DUDE"}

	cards = append(cards, "DU")
	cards.print()
	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }
}

func newCard() string {
	return "5 of Diamonds"
}
