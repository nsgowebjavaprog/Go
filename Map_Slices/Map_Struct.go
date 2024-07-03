package main

import (
	"log"
)

type User struct {
	First_Name string
	Last_Name  string
}

func main() {
	my_map := make(map[string]User)

	me := User{
		First_Name: "Nagaraj",
		Last_Name:  "Loni",
	}

	my_map["me"] = me
	log.Println(my_map["me"].First_Name) //2024/07/03 05:09:25 Nagaraj

}
