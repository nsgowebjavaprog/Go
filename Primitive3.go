package main

import "fmt"

func main() {
	a := 10 // 1010
	b := 3  // 0011

	fmt.Println(a & b)  // 2------>0010
	fmt.Println(a | b)  // 11---->1011
	fmt.Println(a ^ b)  // 9 ---->1001
	fmt.Println(a &^ b) // 8 ----->1000
	fmt.Println(a != b) // true

	fmt.Println(a << 3) // 1010 + 1010 000 ==>16+64 = 80
	fmt.Println(a >> 3) // 1010 - 010 ===> 1 --> 1

	var n complex64 = 1 + 2i
	var s complex64 = 20 + 2i
	fmt.Println(n + s)            // (21+4i)
	fmt.Printf("%v and %T", n, n) // (1+2i) and complex64
	fmt.Printf("%v and %T", s, s)
	fmt.Println()

	s1 := "Nagaraj"
	s2 := " LONI"
	fmt.Printf(s1 + s2) // Nagaraj LONI

	q := "You are useless in our college"
	w := []byte(q)
	fmt.Println(w) // [89 111 117 32 97 114 101 32
	//117 115 101 108 101 115 115 32 105
	// 110 32 111 117 114 32 99 111 108 108 101 103 101]

}
