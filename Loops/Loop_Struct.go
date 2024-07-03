package main

import "fmt"

func main() {
	type User struct {
		First_Name string
		Last_Name  string
		Email      string
		Age        int
	}
	var users []User
	users = append(users, User{"NS", "Loni", "naga.200@gmail.com", 20})

	for _, l := range users {
		fmt.Println(l)
	}
	//{NS Loni naga.200@gmail.com 20}

}
