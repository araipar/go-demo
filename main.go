package main

import "fmt"

type contactInfo struct {
	email string 
	zipCode int
}

type person struct{
	firstName string
	lastName string
	contact contactInfo
}


func main() {

jim := person{
	firstName: "Jim",
	lastName: "Correy",
	contact: contactInfo {
		email:"jim@me.com",
		zipCode:17115,
	},
}

fmt.Printf("%+v",jim)

}


