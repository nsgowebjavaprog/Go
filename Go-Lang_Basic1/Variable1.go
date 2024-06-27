package main

import "fmt"

func main() {
	/*
			var a, b, c, d int = 1, 2, 3, 4

			fmt.Println(a)
			fmt.Println(b)
			fmt.Println(c)
			fmt.Println(d)
		// 1
		// 2
		// 3
		// 4
	*/
	// 2
	/*
		var a, b = 6, "Hello"
		c, d := 7, "NSLONI"

		fmt.Println(a)
		fmt.Println(c)
		fmt.Print(b, " ")
		fmt.Println(d)

		6
		7
		Hello NSLONI

	*/

	var (
		a int
		b int    = 1
		c string = "hello"
	)

	fmt.Println(a) // 0
	fmt.Println(b) // 1
	fmt.Println(c) // hello
}
