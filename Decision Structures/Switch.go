package main

import "fmt"

func main() {
	mt_value := "91"

	switch mt_value {
	case "1":
		fmt.Println("No.1") //No.1
	case "2":
		fmt.Println("No.2")
	case "3":
		fmt.Println("No.3")
	case "4":
		fmt.Println("No.4")
	case "5":
		fmt.Println("No.5")

	default:
		fmt.Println("Default Try Again")
	}
}
