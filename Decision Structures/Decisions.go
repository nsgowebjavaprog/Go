package main

import "log"

func main() {
	my_num := 100
	isTrue := false

	if my_num > 99 && !isTrue {
		log.Println("My number is greater than 90 and Not a true")
		// 2024/07/03 05:35:19 My number is greater than 90 and Not a true
	} else if my_num < 100 && isTrue {
		log.Println("Passing")
	} else if my_num > 1000 && isTrue == false {
		log.Println("Good Passing")
	}
}
