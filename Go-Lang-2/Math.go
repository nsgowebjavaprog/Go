package main

import (
	"fmt"
)

var pf = fmt.Println

func main() { // +,-*,/,%

	pf("3*3:", 3*3) // 3*3: 9

	min_int := 1
	min_int++
	pf(min_int) // 2
}
