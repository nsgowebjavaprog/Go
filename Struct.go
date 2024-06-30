package main

import "fmt"

type Doctor struct {
	number     int
	Doc_name   string
	companions []string
}

func main() {

	a_Doctor := Doctor{
		number:   0000001,
		Doc_name: "Loni_Sir",
		companions: []string{
			"NSLONI",
			"SSLONI",
			"GSLONI",
		},
	}
	fmt.Println(a_Doctor) //{1 Loni_Sir [NSLONI SSLONI GSLONI]}

	fmt.Println(a_Doctor.number) //1
}
