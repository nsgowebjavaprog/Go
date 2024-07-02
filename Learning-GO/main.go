package main

import "fmt"

func main() {
	fmt.Println("Hello, Google")

	var What_you_say string
	var i int

	What_you_say = "Good Morning NS Loni"
	fmt.Println(What_you_say)

	i = 17
	fmt.Println(i)

	result := Say_Some_Thing()
	fmt.Println("Function Returns:", result)

}
func Say_Some_Thing() string {
	return "Bye Nagaraj To See Tommorow"
}

// OUTPUT:
/*
Hello, Google
Good Morning NS Loni
17
Function Returns: Bye Nagaraj To See Tommorow
*/
