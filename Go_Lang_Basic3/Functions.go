package main

import "fmt"

/*
func main() {
	fmt.Println("Wel-Come to Functions")
	greeting1()
	greeting2()
}
func greeting1() {
	fmt.Println("Namstey Go-Lang")
}
func greeting2() {
	fmt.Println("Namstey Go-Lang")
}

Wel-Come to Functions
Namstey Go-Lang
Namstey Go-Lang
*/
// WHEN YOU KNOW no.of element
/*
2
func main() {
	result := adder(3, 6)
	fmt.Println("Result is :", result)
}
func adder(val_one int, val_two int) int {
	return val_one + val_two
}
*/

// Don't Know About No.of Elements

func main() {
	result := n_no_of_ele_adder(2, 3, 5, 6, 4, 3, 5, 3, 53, 5)
	fmt.Println("Result is:", result)
}

func n_no_of_ele_adder(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total // Result is: 89
}
