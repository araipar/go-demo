package main

import "fmt"

type person struct{
	firstName string
	lastName string
}


func main() {

//1.	p := person{"Rai","Paramartha"} // define first string = first attribute , second = second etc
p := person{firstName:"Rai",lastName:"Paramartha"} // 2nd way
fmt.Println(p.firstName)
fmt.Println(p.lastName)
fmt.Println(p)
}


