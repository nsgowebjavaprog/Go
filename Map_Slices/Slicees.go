package main

import (
	"fmt"
	"sort"
)

func main() {
	// creating slice in Go
	var my_slice []int

	my_slice = append(my_slice, 10)
	my_slice = append(my_slice, 50)
	my_slice = append(my_slice, 23)
	my_slice = append(my_slice, 12)
	my_slice = append(my_slice, 29)

	fmt.Println(my_slice) // [10 50 23 12 29]

	sort.Ints(my_slice)
	fmt.Println(my_slice) //[10 12 23 29 50]
}
