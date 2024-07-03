package main

import "log"

func main() {
	var isTrue bool
	isTrue = true

	if isTrue == true {
		log.Println("Is true is: ", isTrue)
		//2024/07/03 05:26:21 Is true is:  true
	} else {
		log.Println("Is False is: ", isTrue)
	}
}
