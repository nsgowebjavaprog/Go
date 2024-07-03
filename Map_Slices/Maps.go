// Map--------> var var_name = map[dt]dt
// key : value

package main

import (
	"fmt"
)

func main() {
	my_map := make(map[string]string)
	my_map["name"] = "Nagaraj"
	my_map["sem"] = "3"
	my_map["college"] = "BiTM Ballary"

	fmt.Println(my_map) // map[college:BiTM Ballary name:Nagaraj sem:3]

	fmt.Println(my_map["name"]) //Nagaraj
}
