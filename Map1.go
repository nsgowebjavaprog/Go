package main

import "fmt"

func main() {
	state_population := make(map[string]int)
	state_population = map[string]int{

		"karnataka": 65434345,
		"Raja":      34345634,
		"Bengal":    23456789,
	}
	fmt.Println(state_population)
	//map[Bengal:23456789 Raja:34345634 karnataka:65434345]
	fmt.Println(state_population["karnataka"]) // 65434345

	delete(state_population, "Raja")
	fmt.Println(state_population)
	// map[Bengal:23456789 karnataka:65434345]

	fmt.Println(state_population["Raja"]) // 0

	state_population["asasas"] = 345345
	fmt.Println(state_population) // map[Bengal:23456789 asasas:345345 karnataka:65434345]

}
