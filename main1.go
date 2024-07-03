// package main

// import (
// 	"fmt"
// 	"net/http"
// )

// func main() {

// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

// 		n, err := fmt.Fprintf(w, "Hello NS LONI")
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		fmt.Println(fmt.Sprintln("No.of Bytes Written: %d", n))
// 		// No.of Bytes Written: 13

// 	})
// 	_ = http.ListenAndServe(":8080", nil)

// }