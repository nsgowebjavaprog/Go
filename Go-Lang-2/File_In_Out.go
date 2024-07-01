package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// Create a new file
	file, err := os.Create("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Write to the file
	fmt.Fprintf(file, "Hello, World!")

	// Read from the file
	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))

	// Rename the file
	err = os.Rename("test.txt", "new_test.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Delete the file
	err = os.Remove("new_test.txt")
	if err != nil {
		log.Fatal(err)
	}
}
