package main

import (
	"fmt"
	"log"
	"net/http"
)

func mainPage(response http.ResponseWriter, request *http.Request) {
	fmt.Println(request.URL)
}

func main() {
	http.HandleFunc("/", mainPage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
