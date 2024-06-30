package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 910}

	b := a[:]   // [1 2 3 4 5 6 7 8 910]
	c := a[3:]  // [4 5 6 7 8 910]
	d := a[:6]  // [1 2 3 4 5 6]
	e := a[3:6] // [4 5 6]
	//f := a[::-1]

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	//fmt.Println(f)
}
