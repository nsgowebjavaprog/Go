package main

import "fmt"

func main() {

	a := make([]int, 3, 100) // a := []int{}
	fmt.Println(a)
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))

	/*
			[0 0 0]
		Length: 3
		Capacity: 100
	*/

	a = append(a, 1)
	fmt.Println(a)
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))
	/*
			[0 0 0 1]
		Length: 4
		Capacity: 100
	*/
}
