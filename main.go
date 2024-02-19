package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	var headers = ""
	for name, values := range r.Header {
		// Loop over all values for the name.
		for _, value := range values {
			headers += name + value
		}
	}
	log.Print(headers)
	fmt.Fprintf(w, "Hello world!")
}

func main() {
	log.Printf("Starting Server")
	http.HandleFunc("/", homeHandler)
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
