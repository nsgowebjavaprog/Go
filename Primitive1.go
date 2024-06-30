package main

import "fmt"

func main() {
	var res bool = true
	fmt.Printf("%v  and  %T\n", res, res) // true  and  bool

	a := 1 == 1
	fmt.Println(a) // true

	e := 1 != 3
	fmt.Print(e) // true

}
