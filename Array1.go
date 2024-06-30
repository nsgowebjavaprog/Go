package main

import (
	"fmt"
)

func main() {
	/*
	   	// arrr_name := [size] int {1,2,3,4........}
	   //1
	   	marks := [5]int{89, 90, 96}
	   	fmt.Printf("Marks %v", marks) // Marks [89 90 96 0 0]
	   	fmt.Println()

	   //2
	   	score := [...]string{"A+", "B+", "A"}
	   	fmt.Println(score) // [A+ B+ A]

	   //3
	   	var student [3]int
	   	student[0] = 12
	   	student[1] = 16
	   	student[2] = 13
	   	fmt.Println(student) // [12 16 13]
	*/
	// 4
	// 2-D Array

	var student = [4][2]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}
	fmt.Println(student) // [[1 2] [3 4] [5 6] [7 8]]

	// 3-D

	var three_D = [2][2][2]int{{{1, 2}, {3, 4}}, {{5, 6}, {7, 8}}}
	fmt.Println(three_D) // [[[1 2] [3 4]] [[5 6] [7 8]]]
}
