package main

import "log"

type Animal interface {
	Say() string
	No_of_Legs() int
}
type Dog struct {
	Name  string
	Breed string
}
type Cat struct {
	Name        string
	Color       string
	No_of_Teeth int
}

func main() {
	// Dog info
	dog := Dog{
		Name:  "Gmy",
		Breed: "Indor",
	}
	Print_Info(&dog)

	//cat info
	cat := Cat{
		Name:        "Jolly",
		Color:       "Black-white",
		No_of_Teeth: 24,
	}
	Print_Info(&cat)
}

func Print_Info(a Animal) {
	log.Println("This animal says", a.Say(), "and has", a.No_of_Legs(), "legs")

	// 	2024/07/03 06:29:44 This animal says WOOFF and has 4 legs
 	// 	2024/07/03 06:29:44 This animal says Meoooo and has 4 legs
	
}

func (d *Dog) Say() string {
	return "WOOFF"
}

func (d *Dog) No_of_Legs() int {
	return 4
}


func (c *Cat) Say() string {
	return "Meoooo"
}

func (c *Cat) No_of_Legs() int {
	return 4
}
