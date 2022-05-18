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

jim.updateName("Rain")

jim.print()

}
func (p person) print(){
	fmt.Printf("%+v",p)
}

func (pointerToPerson *person) updateName(newFirstName string){ 
	(*pointerToPerson).firstName = newFirstName // * operator = give me the value this memory address is pointing at
}

//  0001       person{firstname:"Jim"...}
// address^           value^

// turn address into value with *address
// turn value into address with   &value


// use pointers to change Value Types like : int,float , string ,bool,structs
// dont worry about pointers with Reference Types like : Slices, maps, channels , pointers , functions

