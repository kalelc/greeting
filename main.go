package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", GreetingServer)
	port := "8080"

	fmt.Println("Server started in " + port)
	fmt.Println("Add resource path to url to test, Example: localhost:8080/Andy")
	fmt.Println("Getting Hello, Andy!")
	http.ListenAndServe(":"+port, nil)
}

func GreetingServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
