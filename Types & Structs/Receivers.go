package main

import "log"

type my_struct struct {
	First_Name string
}

func (m *my_struct) printFirst_Name() string {
	return m.First_Name
}

// printFirst_Name ----> method
//

func main() {
	var my_var1 my_struct
	my_var1.First_Name = "NS"

	var my_var2 my_struct
	my_var2.First_Name = "Loni"

	//       {OR}
	//my_var2 := my_struct{
	//	First_Name: "Loni",
	//}

	log.Println("my_var_1: ", my_var1.printFirst_Name())
	log.Println("my_var_2: ", my_var2.printFirst_Name())
}

/*
2024/07/02 18:17:57 my_var_1:  NS
2024/07/02 18:17:57 my_var_2:  Loni
*/
