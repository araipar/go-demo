package main
import "testing"

// test file is named x_test.go

// func test new deck : code to make sure that a deck is made with x number of CARDS, if it doesnt, log it

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected deck Length of 52 but we got %d", len(d))
	}                                               
}