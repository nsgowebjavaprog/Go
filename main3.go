package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// Home Page Handling

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is HOME page")
}

// About Page Handling

func About(w http.ResponseWriter, r *http.Request) {
	sum := AddValues(10, 10)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is Sum 10 + 10 %d", sum))
}

// AddValuea 2 int's

func AddValues(x, y int) int {
	return x + y
}

func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := dividevalue(100.0, 10.0)
	if err != nil {
		fmt.Fprintf(w, "Cannot Divide By 0")
	}
	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f", 100.0, 10.0, f))
}

func dividevalue(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("Cannot Divide by Zero")
		return 0, err
	}
	result := x / y
	return result, nil
}

// Main Function

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Println(fmt.Sprintf("Starting Applicatiuon on Port: %s ", portNumber))
	_ = http.ListenAndServe(portNumber, nil)

}

// output

//Starting Applicatiuon on Port: :8080
