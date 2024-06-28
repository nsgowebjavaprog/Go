package main

import "fmt"

func main() {
	fmt.Println("Wel-Come To Pointers")
	/*
		var ptr *int
		fmt.Println("pointer is: ", ptr)  // pointer is:  <nil>
	*/

	my_num := 20

	var ptr = &my_num

	fmt.Println("Adr of ptr : ", ptr)  // 0x1898068
	fmt.Println("Adr of ptr : ", *ptr) //  20

	*ptr = *ptr * 2
	fmt.Println("New val is:", my_num) // 40

	*ptr = *ptr + 2
	fmt.Println("New val is:", my_num)
}
