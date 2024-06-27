//  %v	Prints the value in the default format
//  %#v	Prints the value in Go-syntax format
//  %T	Prints the type of the value
//  %%	Prints the % sign

package main

import (
	"fmt"
)

func main() {

	var i = 15.5
	var txt = "Hello Ns loni"

	fmt.Printf("%v\n", i)   // 15.5
	fmt.Printf("%T\n", i)   // float64
	fmt.Printf("%#v\n", i)  // 15.5
	fmt.Printf("%v%%\n", i) // 15.5%

	fmt.Printf("%v\n", txt)   // Hello Ns loni
	fmt.Printf("%v%%\n", txt) // Hello Ns loni%
	fmt.Printf("%#v", txt)    // Hello Ns loni
	fmt.Printf("%T\n", txt)   //"Hello Ns loni"string

}
