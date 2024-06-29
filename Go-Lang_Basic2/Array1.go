package main

import "fmt"

func main() {
	fmt.Println("Wwl-Come To Arrray in GO")

	var Fruits_List [5]string

	Fruits_List[1] = "apple"
	Fruits_List[2] = "Banana"
	Fruits_List[4] = "Peach"

	fmt.Println("Fruits List as follows: ", Fruits_List) // Fruits List as follows:  [ apple Banana  Peach]

	fmt.Println("Fruits List as follows: ", len(Fruits_List)) // 5

	fmt.Println("Fruits List as follows: ", Fruits_List[3])

	fmt.Println("Fruits List as follows: ", Fruits_List[2]) // Banana

	var Veg_List = [3]string{"potato", "Beans"}
	fmt.Println(Veg_List) // [potato Beans ]

}
