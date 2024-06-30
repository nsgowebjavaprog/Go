package main

import (
	"fmt"
)

func main() {
	a_doctor := struct{ name string }{name: "Nagaraj"}
	another_doctor := a_doctor
	another_doctor.name = "Raja"

	fmt.Println(a_doctor)       //{Nagaraj}
	fmt.Println(another_doctor) //{Raja}
}
