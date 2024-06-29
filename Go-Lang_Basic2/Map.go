package main

import "fmt"

func main() {
	fmt.Println("Wel-Come to Map in Go-Lang")

	languages := make(map[string]string)

	languages["js"] = "JavaScript"
	languages["ts"] = "TypeScript"
	languages["rb"] = "Ruby"
	languages["rs"] = "Rust"
	languages["go"] = "Go-Lang"
	fmt.Println("List of All Languages:", languages)

	fmt.Println("JS means: ", languages["js"])
	// JS means:  JavaScript

	delete(languages, "rb")
	fmt.Println("Delete ruby from my list: ", languages)
	// Delete ruby from my list:  map[go:Go-Lang js:JavaScript rs:Rust ts:TypeScript]

	// LOooP

	for key, value := range languages {
		fmt.Printf("FOR key %v, value is %v\n", key, value)
	}
	// output:
	/*
		FOR key js, value is JavaScript
		FOR key ts, value is TypeScript
		FOR key rs, value is Rust
		FOR key go, value is Go-Lang
	*/

}
