package main

import "fmt"

func main() {
	switch 2 {
	case 1:
		fmt.Println("Hello,--->1")
	case 2:
		fmt.Println("Hello,--->2") //Hello,--->2
	default:
		fmt.Println("Not MAtching")
	}
}
