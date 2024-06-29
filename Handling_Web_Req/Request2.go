/*
package main

import (

	"fmt"
	"net/http"

)

const url = "https://lco.dev"

	func main() {
		fmt.Println("LCO Web Request")

		response, err := http.Get(url)

		if err != nil {
			panic(err)
		}
		fmt.Println("Respones is of Type : %T\n", response)

		defer response.Body.Close()

}
*/
package main

import (
	"fmt"
	"net/http"
)

const url = "https://www.google.com"

func main() {
	fmt.Println("LCO Web Request")

	response, err := http.Get(url)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println("Response is of Type : %T\n", response)

	defer response.Body.Close()
}
