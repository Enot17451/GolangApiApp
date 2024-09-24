package main

import (
	"html/template"
	"log"
	"net/http"
)

type User struct {
	name string
	age  int
}

func mainPage(response http.ResponseWriter, request *http.Request) {
	user := User{"Mike", 12}
	templ, err := template.ParseFiles("html/main.html")
	if err != nil {
		return
	}
	err = templ.Execute(response, user)
	if err != nil {
		return
	}
}

func main() {
	http.HandleFunc("/", mainPage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
