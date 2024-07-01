package main

import "fmt"

var pf = fmt.Println

type Animal interface {
	Angry_Sound()
	Happy_Sound()
}

type Cat string

func (c Cat) Attack() {
	pf("Cat Attacks It's Prey")
}
func (c Cat) Name() string {
	return string(c)
}

func (c Cat) Angry_Sound() {
	pf("Cat Say HIIIIISSSSSSSSSSSS")
}

func (c Cat) Happy_Sound() {
	pf("Cat Say Meooooooooooo")
}

func main() {
	var kitty Animal
	kitty = Cat("Kitty")
	kitty.Angry_Sound() //Cat Say HIIIIISSSSSSSSSSSS
	kitty.Happy_Sound() //Cat Say Meooooooooooo

}
