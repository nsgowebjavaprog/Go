package main

import "fmt"

var pf = fmt.Println

func Say_Hello() {
	pf("Hello Nagaraj Loni Good Morning")
}

func add(a int, b int) int {
	return a + b //30
}

func main() {
	Say_Hello() //Hello Nagaraj Loni Good Morning
	pf(add(10, 20))
}
