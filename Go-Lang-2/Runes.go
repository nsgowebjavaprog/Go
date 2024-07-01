package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	user_name := "NAGARAJ_nagaraj"
	fmt.Println("Runes Count: ", utf8.RuneCountInString(user_name))

	for i, runes := range user_name {
		fmt.Printf("%d : %#U : %c\n", i, runes, runes)
	}
}

/*
Runes Count:  15
0 : U+004E 'N' : N
1 : U+0041 'A' : A
2 : U+0047 'G' : G
3 : U+0041 'A' : A
4 : U+0052 'R' : R
5 : U+0041 'A' : A
6 : U+004A 'J' : J
7 : U+005F '_' : _
8 : U+006E 'n' : n
9 : U+0061 'a' : a
10 : U+0067 'g' : g
11 : U+0061 'a' : a
12 : U+0072 'r' : r
13 : U+0061 'a' : a
14 : U+006A 'j' : j
*/
