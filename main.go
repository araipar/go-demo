package main

import "fmt"

type person struct{
	firstName string
	lastName string
}


func main() {

var p person
p.firstName = "Jessi"
p.lastName = "Paramartha"
fmt.Println(p)
fmt.Printf("%+v",p) //  print every attribute value

}


