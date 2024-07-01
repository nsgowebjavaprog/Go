package main

import (
	"fmt"
)

func main() {

	Age := 20 //Wow, degree Pass-Out Strudent

	if (Age >= 1) && (Age <= 18) {
		fmt.Println("Child")
	} else if (Age == 20) || (Age == 21) {
		fmt.Println("Wow, degree Pass-Out Strudent")
	} else if Age > 65 {
		fmt.Println("Oh You are Old")
	} else {
		fmt.Println("Try again YouLess")
	}
}
