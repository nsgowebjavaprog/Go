package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Wel-Come to time Study of Go-Lang")

	present_time := time.Now()
	fmt.Println(present_time) // 2024-06-28 07:52:04.263454 +0530 IST m=+0.004586701

	fmt.Println(present_time.Format("01-02-2004 monday")) //[not correct] 06-28-28054 monday

}
