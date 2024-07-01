package main

import (
	"fmt"
)

var pf = fmt.Println

func main() {
	nums := []int{1, 2, 3}
	for _, num := range nums {
		pf(num)
	}

}

/*
1
2
3
*/
