// Golang program to read and write the files
package main

// importing the requires packages
import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func CreateFile(filename, text string) {

	// fmt package implements formatted I/O and
	// contains inbuilt methods like Printf
	// and Scanf
	fmt.Printf("Writing to a file in Go lang\n")

	// Creating the file using Create() method
	// with user inputted filename and err
	// variable catches any error thrown
	file, err := os.Create(filename)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	// closing the running file after the main
	// method has completed execution and
	// the writing to the file is complete
	defer file.Close()

	// writing data to the file using
	// WriteString() method and the
	// length of the string is stored
	// in len variable
	len, err := file.WriteString(text)
	if err != nil {
		log.Fatalf("failed writing to file: %s", err)
	}

	fmt.Printf("\nFile Name: %s", file.Name())
	fmt.Printf("\nLength: %d bytes", len)
}

func ReadFile(filename string) {

	fmt.Printf("\n\nReading a file in Go lang\n")

	// file is read using ReadFile()
	// method of ioutil package
	data, err := ioutil.ReadFile(filename)

	// in case of an error the error
	// statement is printed and
	// program is stopped
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}
	fmt.Printf("\nFile Name: %s", filename)
	fmt.Printf("\nSize: %d bytes", len(data))
	fmt.Printf("\nData: %s", data)

}

// main function
func main() {

	// user input for filename
	fmt.Println("Enter filename: ")
	var filename string
	fmt.Scanln(&filename)

	// user input for file content
	fmt.Println("Enter text: ")
	inputReader := bufio.NewReader(os.Stdin)
	input, _ := inputReader.ReadString('\n')

	// file is created and read
	CreateFile(filename, input)
	ReadFile(filename)
}

/*
PS E:\Go-Lang\Go-Lang-2> go run File_I_O.go
Enter filename:
ns.txt
Enter text:
I am Nagaraj Loni Here i am wring file text to writing and then reading thr text file name of file is ns.txt
Writing to a file in Go lang

File Name: ns.txt
Length: 110 bytes

Reading a file in Go lang

File Name: ns.txt
Size: 110 bytes
Data: I am Nagaraj Loni Here i am wring file text to writing and then reading thr text file name of file is ns.txt
*/
