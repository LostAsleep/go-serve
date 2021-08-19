package main

import (
	"fmt"
	"log"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	rmet := fmt.Sprintf("RequestMethod: %v\n", r.Method)
	ruri := fmt.Sprintf("RequestURI: %v\n", r.RequestURI)

	// Print w/ timestampt to stdout
	log.Printf(ruri)
	log.Printf(rmet)

	//Print to the browser (on served site)
	// fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	fmt.Fprintf(w, ruri)

	//println("Someone just accessed the server.")
	//fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func main() {
	http.HandleFunc("/", sayHello)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
