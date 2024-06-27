package main

import "fmt"

/*
const PI = 3.142
const name = "Nagaraj"

func main() {
	fmt.Println(PI)   // 3.142
	fmt.Println(name) // Nagaraj

}
*/

const (
	A int = 1
	B     = 3.132
	C     = "HI"
)

func main() {

	var i string = "Hello"

	fmt.Printf("i has value: %v and type: %T\n", i, i) //i has value: Hello and type: string

	fmt.Println(A) // 1
	fmt.Println(B) // 3.132
	fmt.Println(C) // HI
}
