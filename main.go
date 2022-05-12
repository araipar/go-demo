package main

//import "fmt"


func main() {


	cards := newDeckFromFile("my_cards")
	cards.print()
	
	//fmt.Println(cards.toString())


}

func newCard() string {
	return "Five of Diamonds"
}
