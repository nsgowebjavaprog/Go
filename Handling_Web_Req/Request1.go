package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Wel-Come To File in Go-Lang")

	content := "This is Inforamation about The Persion  like Who Are Learning Go-Lang1:ABCD2:CDSF3:ERTY4:ERER5:FDHD"

	file, err := os.Create("./mylocalfile.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}

	fmt.Println("Length is: ", length) // Length is:  99

	defer file.Close()
}
