package main

import "fmt"

func main() {

	var i = 15
	fmt.Printf("%b\n", i)  // Binary --> 1111
	fmt.Printf("%d\n", i)  // Decimal ---> 15
	fmt.Printf("%+d\n", i) // +15
	fmt.Printf("%o\n", i)  // 17

	fmt.Printf("%O\n", i)  // 0o17
	fmt.Printf("%x\n", i)  // f
	fmt.Printf("%X\n", i)  // F
	fmt.Printf("%#x\n", i) // 0xf

	fmt.Printf("%4d\n", i)  // __15
	fmt.Printf("%-4d\n", i) // 15
	fmt.Printf("%04d\n", i) // 0015

	var j = true
	var k = false

	fmt.Printf("%t\n", j) // true
	fmt.Printf("%t\n", k) // false
}
