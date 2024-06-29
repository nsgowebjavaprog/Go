// Arrays are used to store multiple values of the same type in a
// single variable, instead of declaring separate variables for
// each value.

// Declare an Array

// var array_name = [length]datatype{values}
//                 {or}
// array_name := [length]datatype{values}

// 1. Eg

package main

import "fmt"

func main() {
	var arr_name1 = [3]int{1, 2, 3}

	// {or}

	arr_name2 := [3]int{1, 2, 3}

	fmt.Println(arr_name1) // [1 2 3]
	fmt.Println(arr_name2) // [1 2 3]

	strs := [4]string{"asas", "asasasas", "asasswdsd", "reewdwe"}
	fmt.Println(strs) // [asas asasasas asasswdsd reewdwe]

	prics := [10]int{12, 12, 23, 3454, 6576, 34, 23254, 464} // idx start
	fmt.Println(prics[2])                                    // 23
	fmt.Println(prics[4])                                    // 6576
	fmt.Println(prics[5])                                    //34
	fmt.Println(prics[6])                                    // 23254

	//8
	prics[8] = 54
	fmt.Println(prics) // [12 12 23 3454 6576 34 23254 464 54 0]
}
