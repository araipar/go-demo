package main


func main() {


	cards := newDeck()
	
	hand, remainingDeck := deal(cards, 5) //  2 new variable in order to get multi value returnm function

	hand.print()
	remainingDeck.print()


	//cards.print()

}

func newCard() string {
	return "Five of Diamonds"
}
