package main

import "fmt"

type deck []string

func (d deck) print() { //any variable can ge access to this method bcos of receiver
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) { //multiplpe value returnin{
	return d[:handSize], d[handSize:]
}
