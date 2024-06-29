package main

import "fmt"

func main() {

	fmt.Println("Structs in Go-Lang")

	loni := User{"Nagaraj", "Back_End@Go.Com", true, 20}
	fmt.Println(loni)
	fmt.Println("Loni Details as Follow:\n", loni)
	fmt.Println("Name %v and Email: %v", loni.Name, loni.Email)
	//  Structs in Go-Lang
	//  {Nagaraj Back_End@Go.Com true 20}
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
