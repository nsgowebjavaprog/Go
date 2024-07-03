package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasGF     bool   `json:"have_gf"`
}

func main() {
	myJson := `
	[
		{
			"first_name":"Ns",
			"last_name":"loni",
			"hair_color":"black",
			"have_gf":true
		},
		{
			"first_name":"SG",
			"last_name":"loni",
			"hair_color":"white",
			"have_gf":false
		}
	]`
	var unmarshalled []Person
	err := json.Unmarshal([]byte(myJson), &unmarshalled)

	if err != nil {
		log.Println("Error unmarshilling json", err)
	}
	log.Println("unmarshalled : %v", unmarshalled)

	// write json from a struct

	var mySlice []Person

	var m1 Person
	m1.FirstName = "NS"
	m1.LastName = "loni"
	m1.HairColor = "black"
	m1.HasGF = true

	mySlice = append(mySlice, m1)

	var m2 Person
	m1.FirstName = "SG"
	m1.LastName = "loni"
	m1.HairColor = "white"
	m1.HasGF = false

	mySlice = append(mySlice, m2)

	newJson, err := json.MarshalIndent(mySlice, "", "    ")

	if err != nil {
		log.Println("Error Marshalling", err)
	}
	fmt.Println(string(newJson))
}

/*
2024/07/03 07:54:34 unmarshalled : %v [{Ns loni black true} {SG loni white false}]
[
    {
        "first_name": "NS",
        "last_name": "loni",
        "hair_color": "black",
        "have_gf": true
    },
    {
        "first_name": "",
        "last_name": "",
        "hair_color": "",
        "have_gf": false
    }
]
*/
