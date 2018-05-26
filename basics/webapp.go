package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Handler - match url <-> function
	// nul = null, none
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/about/", about_handler)
	http.ListenAndServe(":8000", nil)
}

func index_handler(writer http.ResponseWriter, r *http.Request) {
	// Print to writer, not stdout
	fmt.Fprintf(writer, "Go is neat..!")
}

func about_handler(writer http.ResponseWriter, r *http.Request) {
	// Print to writer, not stdout
	fmt.Fprintf(writer, "About Go..?")
	fmt.Fprintf(writer, "Go is awesome!!")
}
