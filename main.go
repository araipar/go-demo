package main

//import "fmt"


func main() {


	cards := newDeck()
	cards.shuffle()
	cards.print()
	
	//fmt.Println(cards.toString())


}

func newCard() string {
	return "Five of Diamonds"
}
