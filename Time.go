package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()

	fmt.Println(now.Year(), now.Month(), now.Day())
	fmt.Println(now.Hour(), now.Minute(), now.Second())
}

/*
2024 June 30
19 50 19
*/
