package main

import (
	"log"
	"time"
)

type User struct {
	First_Name    string
	Last_Name     string
	Phone_Number  string
	Age           int
	Birth_of_Date time.Time
}

func main() {
	log.Println("Hello, NS")
	user := User{
		First_Name:   "Nagaraj",
		Last_Name:    "Loni",
		Phone_Number: "6767676767",
	}
	log.Println(user.First_Name)
	log.Println(user.Birth_of_Date) //2024/07/02 17:51:01 0001-01-01 00:00:00 +0000 UTC

	/*
		2024/07/02 17:49:28 Hello, NS
		2024/07/02 17:49:28 Nagaraj
	*/
}
func some_thing() {
	log.Println("Hello")
}

// When function Nama  is start with Capital Leter --->It visible outside the package
