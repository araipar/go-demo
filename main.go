package main

import "fmt"

type contactInfo struct {
	email string 
	zipCode int
}

type person struct{
	firstName string
	lastName string
	contactInfo
}

func (p person) print(){
	fmt.Printf("%+v",p)
}

func (p person) updateName(newFirstName string){
	p.firstName = newFirstName
}

func main() {

jim := person{
	firstName: "Jim",
	lastName: "Correy",
	contactInfo: contactInfo {
		email:"jim@me.com",
		zipCode:17115,
	},
}

jim.updateName("Rai")
jim.print()

}


