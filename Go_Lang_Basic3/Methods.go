package main

import "fmt"

func main() {

	fmt.Println("Structs in Go-Lang")

	loni := User{"Nagaraj", "Back_End@Go.Com", true, 20}
	fmt.Println(loni)
	fmt.Println("Loni Details as Follow:\n", loni)
	fmt.Println("Name %v and Email: %v", loni.Name, loni.Email)

	loni.Get_Status()

	loni.NewMail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) Get_Status() {
	fmt.Println("Is User Activw: ", u.Status) // Is User Activw:  true
}

func (u User) NewMail() {
	u.Email = "nagarajloni@gmail.com"
	fmt.Println("Email of This User is:", u.Email) //Email of This User is: nagarajloni@gmail.com
}
