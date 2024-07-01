package main

import (
	"fmt"
)

var pf = fmt.Println

type My_Constraint interface {
	int | float64
}

//Defining a Generic Function:

func get_sum_gen[T My_Constraint](x T, y T) T {
	return x + y
}

func main() {
	pf(" 5 + 4 :", get_sum_gen(5, 4))               //9
	pf(" 23.45 + 6.55 :", get_sum_gen(23.45, 6.55)) //30
}
