package main

import "fmt"

var pf = fmt.Println

func fact_num(num int) int {
	if num == 0 {
		return 1
	}
	return num * fact_num(num-1)
}

func main() {
	pf("Factorial of Num 4:", fact_num(4))
	// Factorial of Num 4: 24
	pf("Factorial of Num 5:", fact_num(5))
	// Factorial of Num 5: 120
	pf("Factorial of Num 6:", fact_num(6))
	// Factorial of Num 6: 720
}
