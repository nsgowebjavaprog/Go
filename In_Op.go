package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var pf = fmt.Println

func main() {
	pf("Hell, What is Your Name: ")
	// Taking Input from User
	reader := bufio.NewReader(os.Stdin)

	name, err := reader.ReadString('\n')

	if err == nil {
		pf("Hello", name)
	} else {
		log.Fatal(err)
	}
}
