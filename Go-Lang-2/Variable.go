package main

import (
	"fmt"
	"reflect"
)

var pf = fmt.Println

func main() {
	pf(reflect.TypeOf(27))      // Output: int
	pf(reflect.TypeOf(3.34))    // Output: float64
	pf(reflect.TypeOf(true))    // Output: bool
	pf(reflect.TypeOf("Hello")) // Output: string
	pf(reflect.TypeOf('@'))     // Output: int32
}
