package main

import (
	"log"
)

func main() {
	var my_string string
	my_string = "Nagaraj"

	log.Println("Name of Student Learning Here Go :", my_string)
	//Nagaraj

	student_learning(&my_string)

	log.Println("Name of Student Learing Go With Web Adr's:", my_string)
	//Nagaraj Loni
}
func student_learning(s *string) {

	log.Println("Name of Student Learing Go With Web :", s) //0x1c22060

	new_name := "Nagaraj Loni"
	*s = new_name
}

/*
2024/07/02 17:28:59 Name of Student Learning Here Go : Nagaraj
2024/07/02 17:28:59 Name of Student Learing Go With Web : 0x1c22060
2024/07/02 17:28:59 Name of Student Learing Go With Web Adr's: Nagaraj Loni
*/
