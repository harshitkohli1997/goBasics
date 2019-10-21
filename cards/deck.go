package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type deck []string

func (d deck) print() { //any variable of type deck can get access to this method bcos of receiver
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func deal(d deck, handSize int) (deck, deck) { //multiplpe value returnin{
	return d[:handSize], d[handSize:]
}

func newDeck(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("ERROR", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}
