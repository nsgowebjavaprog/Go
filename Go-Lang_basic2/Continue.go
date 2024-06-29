package main

import "fmt"

func main() {

	val := 1
	for val < 10 {

		if val == 5 {
			val++
			continue
		}
		fmt.Println("Value is : ", val)
		val++
	}
}

/*
Value is :  1
Value is :  2
Value is :  3
Value is :  4
Value is :  6
Value is :  7
Value is :  8
Value is :  9
*/
