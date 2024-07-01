package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type ToDoList struct {
	ToDoCount int
	ToDos     []string
}

func error_check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func write(writer http.ResponseWriter, msg string) {
	_, err := writer.Write([]byte(msg))
	error_check(err)
}

func getStrings(filename string) []string {
	var lines []string
	file, err := os.Open(filename)
	if os.IsNotExist(err) {
		return nil
	}
	error_check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	error_check(scanner.Err())
	return lines
}

func english_handler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Hello Internet")
}

func spanish_handler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Hi Internet")
}

func interact_handler(writer http.ResponseWriter, request *http.Request) {
	todo_vals := getStrings("WebApp.txt")
	fmt.Printf("%#v\n", todo_vals)
	// html
	tmp1, err := template.ParseFiles("View.html")
	error_check(err)
	todos := ToDoList{
		ToDoCount: len(todo_vals),
		ToDos:     todo_vals,
	}
	err = tmp1.Execute(writer, todos)
	error_check(err)
}

func new_handler(writer http.ResponseWriter, request *http.Request) {
	tmp1, err := template.ParseFiles("WebApp.html")
	error_check(err)
	err = tmp1.Execute(writer, nil)
	error_check(err)
}

func create_handler(writer http.ResponseWriter, request *http.Request) {
	todo := request.FormValue("todo")
	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	file, err := os.OpenFile("WebApp.txt", options, os.FileMode(0600))
	error_check(err)
	_, err = fmt.Fprintln(file, todo)
	error_check(err)
	err = file.Close()
	error_check(err)
	http.Redirect(writer, request, "/interact", http.StatusFound)
}

func main() {
	http.HandleFunc("/Hello", english_handler)
	http.HandleFunc("/Hi", spanish_handler)
	http.HandleFunc("/interact", interact_handler)
	http.HandleFunc("/new", new_handler)
	http.HandleFunc("/create", create_handler)
	http.ListenAndServe(":8080", nil)
}
