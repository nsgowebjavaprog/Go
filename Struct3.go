package main

import "fmt"

type Animal struct {
	Name   string
	Origin string
}

type Bird struct {
	Animal
	Speed   float32
	Can_Fly bool
}

func main() {
	b := Bird{
		Animal:  Animal{Name: "Emu", Origin: "India"},
		Speed:   45.74,
		Can_Fly: false,
	}
	fmt.Println(b.Animal.Name) //Emu

	fmt.Println(b.Can_Fly) //false
	//fmt.Println(b.Animal.Speed)

	fmt.Println(b.Animal.Origin) //India

}
