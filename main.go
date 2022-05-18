package main

import "fmt"


func main() {

	// colors := map[string]string{
	// 	"red" : "#ff0000",
	// 	"green" : "#1234556",
	// }

	colors := make(map[string]string)
	colors["white"] = "#ffffff" // add

	delete(colors,"white") // delete

	fmt.Println(colors)

}
// go is a pass by value language
