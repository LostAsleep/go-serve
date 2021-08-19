package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	// Print w/ timestampt to stdout
	log.Printf("I see you used Method %q\n", r.Method)
	//Print to the browser (on served site)
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))

	//println("Someone just accessed the server.")
	//fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func main() {
	http.HandleFunc("/", sayHello)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
