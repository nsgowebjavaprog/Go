package main

import "fmt"

func change_val(myPtr *int) {
	*myPtr = 10
}

func main() {
	val_1 := 100
	fmt.Println("Before Value is:", val_1)
	change_val(&val_1)
	fmt.Println("After Value is:", val_1)
}

/*
Before Value is: 100
After Value is: 10
*/
