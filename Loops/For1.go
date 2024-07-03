package main

import "fmt"

func main() {

	animals := []string{"Cat", "Dog", "Cow", "Ox", "Horse"}

	for i, animals := range animals {
		fmt.Println(i, animals)
		/*
			0 Cat
			1 Dog
			2 Cow
			3 Ox
			4 Horse
		*/
	}
}
