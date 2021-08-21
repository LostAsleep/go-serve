package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
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

func serveFile(w http.ResponseWriter, r *http.Request) {
	//http.ServeFile(w, r, r.URL.Path)  // Exposes whole file system!
	var err error
	wd, err := os.Getwd()
	if err != nil {
		log.Print(err)
		return
	}
	// GetWD and Join used to create a safe absolute path to serve.
	// For some reason this will also make the logo.png work.
	// With (w, r, ".") the logo.png would not show up.
	http.ServeFile(w, r, filepath.Join(wd, r.URL.Path))
}

func main() {
	http.HandleFunc("/form", serveForm)
	http.HandleFunc("/", serveFile)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
