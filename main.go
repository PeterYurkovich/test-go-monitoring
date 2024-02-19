package main

import (
	"fmt"
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	var headers = ""
	for name, values := range r.Header {
		// Loop over all values for the name.
		for _, value := range values {
			headers += name + value
		}
	}
	log.Printf(headers)
	fmt.Fprintf(w, "Hello world!")
}

func main() {
	log.Printf("Starting Server")
	http.HandleFunc("/", homeHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
