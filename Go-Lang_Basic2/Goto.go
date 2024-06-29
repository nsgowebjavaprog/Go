package main

import "fmt"

func main() {

	val := 1
	for val < 10 {

		if val == 5 {
			goto nsloni
		}
		fmt.Println("Value is : ", val)
		val++
	}

nsloni:
	fmt.Println("wel-Come To Go-Lang Play_List")
}

/*
Value is :  1
Value is :  2
Value is :  3
Value is :  4
wel-Come To Go-Lang Play_List
*/
