package main

import (
	"fmt"
	"log"
	"net/http"
)

func returnForm(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // PostForm is only available after ParseForm is called.
	ruri := fmt.Sprintf(`
	RequestURI: %v
	Host:       %v
	Values:     %v
	`, r.RequestURI, r.Host, r.Form)
	//log.Printf(ruri)  // Print w/ timestampt to stdout
	fmt.Fprintf(w, ruri) // Print to the browser (on served site)
}

func main() {
	http.HandleFunc("/", returnForm)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
