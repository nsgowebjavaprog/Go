package main

import "fmt"

func main() {

	fmt.Println("Wel-Come To LOOP/ BREAK/ VCONTINUE/ GOTO")

	days := []string{"sunday", "thesday", "wendnesday", "Friday", "Saturday", "sunday"}
	//   [sunday thesday wendnesday Friday Saturday sunday]
	fmt.Println(days)

	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}
	/*
			sunday
		thesday
		wendnesday
		Friday
		Saturday
		sunday
	*/

	for k := range days {
		fmt.Println(days[k])
	}
	/*
	   sunday
	   sunday
	   thesday
	   wendnesday
	   Friday
	   Saturday
	   sunday
	*/

	for index, day := range days {
		fmt.Printf("index is %d and value is %v\n", index, day)
	}
	/*
			index is 0 and value is sunday
		index is 1 and value is thesday
		index is 2 and value is wendnesday
		index is 3 and value is Friday
		index is 4 and value is Saturday
		index is 5 and value is sunday
	*/
}
