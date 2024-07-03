package main

import "fmt"

func main() {
	var First_Name = "Uu NAGARAJ nagaraj"
	//First_Name = "a" //0 : 97

	for i, l := range First_Name {
		fmt.Println(i, ":", l)
	}
	/*
		0 : 85
		1 : 117
		2 : 32
		3 : 78
		4 : 65
		5 : 71
		6 : 65
		7 : 82
		8 : 65
		9 : 74
		10 : 32
		11 : 110
		12 : 97
		13 : 103
		14 : 97
		15 : 114
		16 : 97
		17 : 106
	*/
}
