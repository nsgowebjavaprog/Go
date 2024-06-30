package main

import "fmt"

const (
	a    = iota + 5 // 0 + 5
	cat             // 5 + 1 = 6
	dog             // 7
	boll            // 8
)

func main() {

	fmt.Printf("%v\n", cat) // 6

}
