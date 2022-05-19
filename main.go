package main
import "fmt"

type bot interface{  // cara  bacanya : tipe baru (bot) untunk handle segala type yang punya method getGreeting() 
	getGreeting() string


}

type englishBot struct{}

type spanishBot struct{} 

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

}

func printGreeting(b bot){
	fmt.Println(b.getGreeting())
}



func (eb englishBot) getGreeting() string{  //method 
	// custom logic for generating english greeting
	return "Hi there"
}


func (sb spanishBot) getGreeting() string{
	// custom logic for generating spanis greeting

	return "Hola"
}


