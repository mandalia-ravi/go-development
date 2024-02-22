package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}

func formHandler(responseWriter http.ResponseWriter, request *http.Request) {
	if err := request.ParseForm(); err != nil {
		fmt.Fprintf(responseWriter, "ParseForm error: %v", err)
		return
	}
	fmt.Fprintf(responseWriter, "POST request successful\n")
	name := request.FormValue("name")
	address := request.FormValue("address")
	fmt.Fprintf(responseWriter, "Name: %s\n", name)
	fmt.Fprintf(responseWriter, "Address: %s\n", address)
}

func helloHandler(responseWriter http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/hello" {
		http.Error(responseWriter, "404 Not Found", http.StatusNotFound)
		return
	}

	if request.Method != "GET" {
		http.Error(responseWriter, "Method not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(responseWriter, "hello!")
}
