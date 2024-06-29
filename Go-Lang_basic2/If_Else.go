package main

import "fmt"

func main() {
	fmt.Println("Wel-Come To If-Else")

	ns_age := 20
	var result string

	if ns_age < 18 {
		result = "Child"
	} else if ns_age > 60 {
		result = "Pencen"
	} else {
		result = " Exactly 20"
	}
	fmt.Println(result) //  Exactly 20

	if num := 11; num < 10 {
		fmt.Println("Less Than 10")
	} else {
		fmt.Println("Num more than 10") // Num more than 10
	}
}
