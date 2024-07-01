package main

import (
	"fmt"
)

var pf = fmt.Println

func main() {
	var arr1 [5]int
	arr1[1] = 10

	arr2 := [5]int{1, 2, 3, 4, 5}
	pf("Index: ", arr2[2]) //Index:  3

	pf("Length is : ", len(arr2)) //Length is :  5

	for i := 0; i < len(arr2); i++ {
		pf(arr2[i])
	}
	/*
	   1
	   2
	   3
	   4
	   5
	*/

	arr3 := [2][2]int{{2, 2}, {3, 3}}
	pf(arr3)

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			pf(arr3[i][j])
		}
	}
	/*
			[[2 2] [3 3]]
		2
		2
		3
		3
	*/
}
