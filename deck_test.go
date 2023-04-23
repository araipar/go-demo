package main
import (
	"testing"
	"os"
)

// test file is named x_test.go

// func test new deck : code to make sure that a deck is made with x number of CARDS, if it doesnt, log it

//test

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected deck Length of 52 but we got %d", len(d))
	}                 
	
	if d[0] != "Ace Of Spades"{
		t.Errorf("Expected Ace Of Spades on top but we got %s", d[0])
	}

	if d[len(d) - 1] != "K Of Diamonds"{
		t.Errorf("Expected K Of Diamonds on bottom but we got %s", d[len(d) - 1])
	}
	
}

func TestSaveToDeckAndTestNewDeckFromFile(t *testing.T){
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile ("_decktesting")
	if len(loadedDeck) != 52 {
		t.Errorf("Expected deck Length of 52 but we got %d", len(loadedDeck))
	}  

	os.Remove("_decktesting")

}