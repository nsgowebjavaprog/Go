package main

import (
	"fmt"
	"time"
)

var pf = fmt.Println

func print_to_15() {
	for i := 1; i <= 15; i++ {
		pf("Function 1:", i)
	}
}

func print_to_10() {
	for i := 1; i <= 10; i++ {
		pf("Function 2:", i)
	}
}

func main() {
	go print_to_15()
	go print_to_10()
	time.Sleep(2 * time.Second)
}

/*
Function 1: 1
Function 1: 2
Function 1: 3
Function 1: 4
Function 1: 5
Function 1: 6
Function 1: 7
Function 1: 8
Function 1: 9
Function 1: 10
Function 1: 11
Function 1: 12
Function 1: 13
Function 1: 14
Function 1: 15

Function 2: 1
Function 2: 2
Function 2: 3
Function 2: 4
Function 2: 5
Function 2: 6
Function 2: 7
Function 2: 8
Function 2: 9
Function 2: 10
*/
