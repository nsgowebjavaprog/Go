package main

import (
	"fmt"
)

var pf = fmt.Println

func main() {
	for i := 0; i < 10; i++ {
		pf(i)
	}
}

/*
0
1
2
3
4
5
6
7
8
9
*/
