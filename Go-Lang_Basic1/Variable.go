package main

import "fmt"

func main() {

	var user_name string = "Nagaraj"
	fmt.Println(user_name)
	fmt.Printf("variable Type: %T \n", user_name) // variable Type: string

	var log_in bool = true
	fmt.Println(log_in)
	fmt.Printf("variable Type: %T \n", log_in) //variable Type: bool

	var small_val uint16 = 256 // unit8{<255} ,16{2**15},32,64
	fmt.Println(small_val)
	fmt.Printf("variable Type: %T \n", small_val) // variable Type: uint16

	// default values and same aliases

	var website = "learncodeline.in"
	fmt.Println(website)
	fmt.Printf("variable Type: %T \n", website) // string

	// No Variable Style

	age := 21
	fmt.Println(age) // 21

	const PI float64 = 3.142
	fmt.Println(PI) // 3.142

}
