package main
import "fmt"
//Create a new type of deck
// which is a slicew of strings
type deck []string

func (d deck) print(){
	for i, card := range d{
		fmt.Println(i,card)
	}
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