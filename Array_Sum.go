package main

import "fmt"

func get_array_sum(arr []int) int {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return sum
}

func main() {
	arr1 := []int{11, 22, 33, 44, 55}
	fmt.Println("Total Array Sum is:", get_array_sum(arr1))
	//Total Array Sum is: 165
}
