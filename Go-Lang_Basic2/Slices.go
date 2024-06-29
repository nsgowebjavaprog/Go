package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Wel-Come To Slices")

	var list = []string{"Apple", "Banana"}
	fmt.Println("Data Type is: %T\n", list)

	list = append(list, "Mango", "Cherry") //[Apple Banana Mango Cherry]
	fmt.Println(list)

	list = append(list[1:])
	fmt.Println(list) // [Banana Mango Cherry]

	list = append(list[1:3])
	fmt.Println(list) // [Mango Cherry]

	/*	list = append(list[0:3])
		fmt.Println(list)
	*/

	High_Score := make([]int, 5)

	High_Score[0] = 100
	High_Score[1] = 200
	High_Score[2] = 300
	High_Score[3] = 400
	High_Score[4] = 500

	fmt.Println(High_Score) // [100 200 300 400 500]

	High_Score = append(High_Score, 111, 222, 333)

	fmt.Println(High_Score) // [100 200 300 400 500 111 222 333]

	// SoRT

	sort.Ints(High_Score)
	fmt.Println(High_Score) // [100 111 200 222 300 333 400 500]

	// REMOVE A VALUE FROM A CLICE BASED ON INDEX IN GO-LANG

	// hOW TO REMOVE VAL SLICES BASED ON INDEX

	var course = []string{"html", "css", "js", "ts", "go", "sveltkit", "rust", "sql", "telwind-css"}
	fmt.Println(course)

	// remove [rust] because not use in this project

	var index int = 6
	course = append(course[:index], course[index+1:]...)
	fmt.Println(course) // [html css js ts go sveltkit sql telwind-css]
}
