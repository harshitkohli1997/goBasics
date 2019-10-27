package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck("myfile")

	if len(d) != 4 {
		t.Errorf("length must be 4")
	}
}
