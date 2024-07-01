package main

import "fmt"

func main() {
	// Create a map
	flowerColor := map[string]string{
		"Sunflower": "Yellow",
		"Jasmine":   "White",
		"Hibiscus":  "Red",
	}

	fmt.Println(flowerColor) //map[Hibiscus:Red Jasmine:White Sunflower:Yellow]
}

/*
package main

import "fmt"

func main() {
    // Create a map
    currencyCode := map[string]string{
        "USD": "US Dollar",
        "GBP": "Pound Sterling",
        "EUR": "Euro",
    }

    fmt.Println("Map before deletion: ", currencyCode)

    // Delete the key EUR
    delete(currencyCode, "EUR")

    fmt.Println("Map after deletion: ", currencyCode)
}
*/
