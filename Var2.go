package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int = 23
	fmt.Printf("%v %T", i, i) //23 int
	fmt.Println()

	var j string
	j = strconv.Itoa(i)
	fmt.Printf("%v and %T", j, j)
	// 23 int
	// 23 and string

}
