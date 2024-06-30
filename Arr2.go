package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3}
	b := a
	b[1] = 5
	fmt.Println(a)                    // [1 2 3]
	fmt.Println("CAPACITY: ", cap(a)) // CAPACITY:  3
	fmt.Println(b)                    // [1 5 3]

}
