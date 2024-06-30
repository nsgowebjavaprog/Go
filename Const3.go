package main

import "fmt"

func main() {
	// const + int  ====> int
	const a = 32
	var b int32 = 34
	fmt.Printf("%v and %T", a+b, a+b) // 66 and int32
}
