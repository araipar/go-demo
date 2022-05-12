package main
import (
	"fmt"
	"strings"
	"os"
	"io/ioutil"
	"math/rand"
)
//Create a new type of deck
// which is a slicew of strings
type deck []string

func (d deck) print(){ // receiver d deck
	for i, card := range d{
		fmt.Println(i,card)
	}
}

func  deal(d deck, handSize int) (deck,deck) { // return multipple value : deck & deck

	return d[:handSize], d[handSize:]

}

func (d deck) toString() string { // receiver d deck 

	return strings.Join([]string(d),",")
	
}

func (d deck) saveToFile(filename string) error{  //receiver deck , parameter string , returns error

return	ioutil.WriteFile(filename,[]byte(d.toString()), 0666)

}

func newDeck() deck {   // return value deck
	cards := deck{}
	cardSuits := []string{"Spades","Hearts","Clovers","Diamonds"}
	cardValues := []string{"Ace","2","3","4","5","6","7","8","9","10","J","Q","K"}

	for _, suit := range cardSuits{
		for _, value := range cardValues{
			cards = append(cards, value +" Of "+ suit)
		}
	}

	return cards

}

func newDeckFromFile(filename string) deck{
	bs, err := ioutil.ReadFile(filename) // err = value of error. but if everything is ok, err will be nil
	if err != nil{ // if something goes wrong
		// opt 1 : log the error , return a call to new deck 
		// opt 2 : log the error , quit program

		fmt.Println("Error : ", err)
		os.Exit(1)
	}

	 s := strings.Split(string(bs),",")
	 return deck(s)
	
}

func (d deck) shuffle() {
	for i := range d {
		newPosition := rand.Intn(len(d) -1)
		
		d[i], d[newPosition] = d[newPosition], d[i]  // fancy one line array index value swapping 

	}

}