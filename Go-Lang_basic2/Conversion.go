package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Wel-Come To Pizza App")
	fmt.Println("Please Rate our Pizza B/W 1 to 5")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks giving  good rating", input)
}
