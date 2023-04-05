package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

// handle echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "URL.PATH = %q\n", r.URL.Path)
	if err != nil {
		log.Println(err)
	}
}
