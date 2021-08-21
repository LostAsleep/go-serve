package main

import (
	"fmt"
	"log"
	"net/http"
)

const tpl = `
RequestURI: %v
Host:       %v
Form:       %v
    some:   %v
`

func serveForm(w http.ResponseWriter, r *http.Request) {
	// PostForm is only available after ParseForm is called.
	r.ParseForm()
	fmt.Fprintf(w, fmt.Sprintf(tpl, // Print to browser
		r.RequestURI,
		r.Host,
		r.Form,
		r.Form.Get("some"),
	))
}

func some(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "something")
}

func main() {
	http.HandleFunc("/", serveForm)
	http.HandleFunc("/some", some)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
