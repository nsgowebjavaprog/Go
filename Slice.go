package main

import (
	"fmt"
)

var pf = fmt.Println

func main() {
	sl_1 := make([]string, 6)
	sl_1[0] = "Arjun Mahindra"
	sl_1[1] = "Sonalik"
	sl_1[2] = "New Holend"
	sl_1[3] = "Swaraj"
	sl_1[4] = "John Deer"
	pf("Slice is: ", sl_1)            //[Arjun Mahindra Sonalik New Holend Swaraj John Deer ]
	pf("Length of Slice:", len(sl_1)) //Length of Slice: 6

	for i := 0; i < len(sl_1); i++ {
		pf(sl_1[i])
		/*
					Arjun Mahindra
			Sonalik
			New Holend
			Swaraj
			John Deer
		*/

		for _, x := range sl_1 {
			pf(x)
		}

		slice_1 := [5]int{10, 20, 30, 40, 50}
		pf(slice_1)

		pf(slice_1[1:3]) //[20 30]
		pf(slice_1[:5])  //all
		pf(slice_1[:])   //all
		slice_1[0] = 100
		pf(slice_1) //[100 20 30 40 50]
	}
}
