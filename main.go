package main

import "fmt"


func main() {

	colors := map[string]string{
		"red" : "#ff0000",
		"green" : "#1234556",
		"white" : "#fffff",
	}



	printMap(colors)

}

func printMap(c map[string]string){
	for color,hex := range c{  // key,value range args param .  vars created must be used below
		fmt.Println("color " , color, " key is " ,hex)
	}
}

// go is a pass by value language


//maps vs structs :

// https://www.udemy.com/course/go-the-complete-developers-guide/learn/lecture/7797370#notes  , 0:23