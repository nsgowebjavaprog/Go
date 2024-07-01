package main

import "fmt"

func getAvg(nums ...float32) float32 {
	var sum float32 = 0.0
	for _, val := range nums {
		sum += val
	}
	return (sum / 3)
}

func main() {
	isslice := []float32{12, 14, 18}
	fmt.Printf("Average is:%.3f\n", getAvg(isslice...))
}

//Average is:14.667
