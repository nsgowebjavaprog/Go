package main

import "log"

var s1 = "Outer___1"

func main() {
	var s2 = "Inner____1"

	log.Println("Outer: ", s1)
	log.Println("Inner: ", s2)

}

/*
2024/07/02 17:40:39 Outer:  Outer___1
2024/07/02 17:40:39 Inner:  Inner____1
*/
