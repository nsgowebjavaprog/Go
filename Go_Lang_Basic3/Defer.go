package main

import "fmt"

func main() {
	fmt.Println("Hello")
	defer fmt.Println("Google")
	defer fmt.Println("Wipro")
	defer fmt.Println("EY")
	defer fmt.Println("CISCO")
	//Hello
	//CISCO
	//EY
	//Wipro
	//Google

}
