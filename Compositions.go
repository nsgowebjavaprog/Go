package main

import "fmt"

var pf = fmt.Println

type contact struct {
	fname string
	lname string
	phone string
}

type business struct {
	name    string
	address string
	contact
}

func (b business) info() {
	fmt.Printf("Contact at : %s is %s %s ", b.name, b.contact.fname, b.contact.lname)

	//Contact at : ABC Planning is NSLONI LONI
}

func main() {
	con_1 := contact{
		"NSLONI",
		"LONI",
		"286112",
	}
	bus_1 := business{
		"ABC Planning",
		"234 North India",
		con_1,
	}
	bus_1.info()
}
