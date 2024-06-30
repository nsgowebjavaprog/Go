package main

import "fmt"

func main() {
	const a int = 12
	const b string = "Nagaraj"
	const c float32 = 3.241
	const d bool = false

	fmt.Printf("%v and %T", a, a)
	fmt.Println()
	fmt.Printf("%v and %T", b, b)
	fmt.Println()
	fmt.Printf("%v and %T", c, c)
	fmt.Println()
	fmt.Printf("%v and %T", d, d)

}

/*
12 and int
Nagaraj and string
3.241 and float32
false and bool
*/
