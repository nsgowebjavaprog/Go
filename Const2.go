package main

import "fmt"

const a float32 = 3.2323

func main() {
	const a int = 12

	fmt.Printf("%v and %T", a, a) // 12 and int

}
