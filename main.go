package main

import (
	"log"
	"net/http"
)

func main() {
	// After this there is literally a server running!
	log.Fatal(http.ListenAndServe(":8080", nil))
}
