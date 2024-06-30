package main

import (
	"fmt"
)

func main() {
	number := 56
	guess := 56

	if guess < number {
		fmt.Println("Gueess is low") //100
	}
	if guess > number {
		fmt.Println("Gueess id High") // 46
	}
	if guess == number {
		fmt.Println("Yes, You WiN") // 56
	}
}
