package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Wel-Come"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the Rating : ")

	// comma ok {or}  error ok

	input, _ := reader.ReadString('\n')
	fmt.Println("Rating is: ", input)
}
