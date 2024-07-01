package main

import (
	"fmt"
)

var pf = fmt.Println

func main() {
	i := 0
	for i < 100000 {
		pf(i)
		i++
	}

}
