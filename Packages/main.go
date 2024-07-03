package main

import (
	helpers "_/E_/Go_Web/Packages"
	"log"
)

func main() {
	log.Println("Hello")

	var my_var helper.SomeTtpe
	my_var.TypeName = "NS LONI"
	log.Println(my_var.TypeName)
}
