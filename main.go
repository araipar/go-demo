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



func main() {

jim := person{
	firstName: "Jim",
	lastName: "Correy",
	contactInfo: contactInfo {
		email:"jim@me.com",
		zipCode:17115,
	},
}
// go is a pass by value language
jimPointer := &jim  //  & operator = give me the memory address of the value this variable is pointing at

jimPointer.updateName("Rai")

jim.print()

}
func (p person) print(){
	fmt.Printf("%+v",p)
}

func (pointerToPerson *person) updateName(newFirstName string){ // * operator = give me the value this memory address is pointing at
	(*pointerToPerson).firstName = newFirstName
}

//  0001       person{firstname:"Jim"...}
// address^           value^

// turn address into value with *address
// turn value into address with   &value

