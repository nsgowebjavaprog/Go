// package main

// import (
// 	"fmt"
// 	"net/http"
// )

// const portNumber = ":8080"

// // Home Page Handling

// func Home(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "This is HOME page")
// }

// // About Page Handling

// func About(w http.ResponseWriter, r *http.Request) {
// 	sum := AddValues(10, 10)
// 	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is Sum 10 + 10 %d", sum))
// }

// // AddValuea 2 int's

// func AddValues(x, y int) int {

// 	return x + y

// }

// // Main Function

// func main() {

// 	http.HandleFunc("/", Home)
// 	http.HandleFunc("/about", About)

// 	fmt.Println(fmt.Sprintf("Starting Applicatiuon on Port: %s ", portNumber))
// 	_ = http.ListenAndServe(portNumber, nil)

// }

// // output

// //Starting Applicatiuon on Port: :8080
